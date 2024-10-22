package handlers

import (
	"db-test/db"
	"db-test/templates"
	"db-test/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type postsInfo struct {
	LoggedIn       bool
	Title          string
	UsernameForNav string
	IdForNav       string
	PgNum          int
	Pages          int
	Posts          []*postsPageInfo
}

type postsPageInfo struct {
	Logged         bool
	PostId         int
	PublisherName  string
	PostTitle      string
	PostBody       string
	PostDate       string
	PostCategories string
	PostLike       int
	PostDislike    int
	PostComments   int
}

func HandlePostsPage(w http.ResponseWriter, r *http.Request) {
	isLogged := true
	cookie, err := utils.CheckIfAuth(r)
	if err != nil {
		isLogged = false
		// http.Redirect(w, r, "/signIn", http.StatusSeeOther)
	}

	// Getting the categories from the user
	r.ParseForm()
	selectedCategories := r.Form["Category"]

	pageNumber, err := strconv.Atoi(r.PathValue("pgNumber"))
	if err != nil {
		fmt.Println("Error parsing page number:", err)
		HandleBadRequest(w, r)
		return
	}

	// Reset page number to 1 if filters are applied
	//   if len(selectedCategories) > 0 {
	//         pageNumber = 2
	//     }

	// Construct query based on selected categories
	query := ""
	if len(selectedCategories) > 0 {
		query = `
        SELECT
            users.username, 
            posts.title, 
            posts.content, 
            posts.created_at, 
            posts.id,
            COUNT(CASE WHEN post_interactions.vote = 0 THEN 1 END) AS upvotes,
            COUNT(CASE WHEN post_interactions.vote = 1 THEN 1 END) AS downvotes
        FROM posts
        JOIN users ON posts.userid = users.id
        LEFT JOIN post_interactions ON posts.id = post_interactions.postid
        WHERE posts.id IN (
            SELECT Cata_post.Post_ID 
            FROM Cata_post 
            WHERE Cata_post.Cata_ID IN (` + strings.Join(selectedCategories, ",") + `)
        )
        GROUP BY posts.id, users.username, posts.title, posts.content, posts.created_at
        ORDER BY posts.created_at DESC;`
	} else {
		query = `SELECT 
            users.username, 
            posts.title, 
            posts.content, 
            posts.created_at, 
            posts.id,
           COUNT(CASE WHEN post_interactions.vote = 0 THEN 1 END) AS upvotes,
            COUNT(CASE WHEN post_interactions.vote = 1 THEN 1 END) AS downvotes
        FROM posts
        JOIN users ON posts.userid = users.id
        LEFT JOIN post_interactions ON posts.id = post_interactions.postid
        GROUP BY posts.id, users.username, posts.title, posts.content, posts.created_at
        ORDER BY posts.created_at DESC`
	}

	checkers := []string{}
	for _, cat := range selectedCategories {
		var check int
		for _, thing := range checkers {
			if thing == cat {
				HandleBadRequest(w, r)
				return
			}
		}
		checkers = append(checkers, cat)
		db.Database.QueryRow("SELECT COUNT(*) FROM category WHERE id = ?", cat).Scan(&check)
		if check == 0 {
			HandleBadRequest(w, r)
			return
		}
	}

	rows, err := db.Database.Query(query)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return
	}
	defer rows.Close()

	arr := []*postsPageInfo{}
	for rows.Next() {
		pid := ""
		categories := ""
		p := postsPageInfo{}
		err := rows.Scan(
			&p.PublisherName,
			&p.PostTitle,
			&p.PostBody,
			&p.PostDate,
			&pid,
			&p.PostLike,
			&p.PostDislike,
		)

		p.Logged = isLogged
		if err != nil {
			log.Printf("Error reading row: %v", err)
			continue
		}

		postid, err := strconv.Atoi(pid)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		p.PostId = postid

		num := noComments(pid)
		p.PostComments = num
		cats, err := db.Database.Query(
			`SELECT category.title FROM Cata_post JOIN category ON Cata_post.Cata_ID = category.id WHERE Cata_post.Post_ID = ?`,
			pid,
		)
		if err != nil {
			fmt.Println("Error querying categories:", err)
			return
		}
		defer cats.Close()

		for cats.Next() {
			temp := ""
			cats.Scan(&temp)
			categories += temp + " "
		}
		p.PostCategories = categories

		arr = append(arr, &p)
	}

	noPosts := len(arr)
	noPgs := (noPosts / 50)
	if noPosts%50 != 0 {
		noPgs++
	}
	fmt.Println(noPgs)
	fmt.Println(len(arr))

	if (pageNumber > noPgs || pageNumber <= 0) && len(arr) != 0 {
		fmt.Printf("Number of pages: %v ", noPgs)
		HandleNotFound(w, r)
		return
	}

	toDisplay := []*postsPageInfo{}
	fmt.Println(pageNumber)

	if len(arr) != 0 {
		if pageNumber != noPgs {

			pageNumber--
			toDisplay = arr[pageNumber*50 : pageNumber*50+50]
		} else {

			pageNumber--
			toDisplay = arr[pageNumber*50:]

		}
	}

	if isLogged {
		user, err := db.ReadUser(cookie.UserID)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			HandleInternalServerError(w, r)
			return
		}
		userid := strconv.Itoa(user.ID)
		pageNumber++
		err = templates.LoadPage(
			w,
			"pages/posts.html",
			postsInfo{
				LoggedIn:       isLogged,
				Title:          "Home",
				IdForNav:       userid,
				UsernameForNav: user.Username,
				Posts:          toDisplay,
				PgNum:          pageNumber,
				Pages:          noPgs,
			},
		)
		if err != nil {
			fmt.Printf("Error loading page: %v\n", err)
			templates.LoadPage(w, "pages/error.html", err)
			return
		}
	} else {
		pageNumber++
		err = templates.LoadPage(w, "pages/posts.html", postsInfo{LoggedIn: isLogged, Title: "Home", Posts: toDisplay, PgNum: pageNumber, Pages: noPgs})
		if err != nil {
			fmt.Printf("Error loading page: %v\n", err)
			templates.LoadPage(w, "pages/error.html", err)
			return
		}
	}
}
