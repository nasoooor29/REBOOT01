package handlers

import (
	"db-test/db"
	"db-test/templates"
	"db-test/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type info struct {
	LoggedIn       bool
	Title          string
	UsernameForNav string
	IdForNav       string
	ArrInfo        []*postInfo
}

type postInfo struct {
	PublisherName string
	PostTitle     string
	PostBody      string
	PostDate      string
	PostComments  int
}

// not sure where keep this :)
func noComments(postID string) int {
	ans, err := db.Database.Query(`SELECT id FROM comments WHERE postid = ?`, postID)
	if err != nil {
		return -1
	}

	count := 0
	for ans.Next() {
		count++
	}

	return count
}

func HandleHomePage(w http.ResponseWriter, r *http.Request) {
	username := ""
	id := ""
	isLogged := true
	if r.URL.Path != "/" {
		HandleNotFound(w, r)
		return
	}

	cookie, err := utils.CheckIfAuth(r)
	if err != nil {
		isLogged = false
	}

	if isLogged {
		mdl, err := db.ReadUser(cookie.UserID)
		if err != nil {
			//fmt.Println("MEOWZER")
		}
		username = mdl.Username
		id = strconv.Itoa(mdl.ID)
	}

	rows, err := db.Database.Query(`SELECT users.username, posts.title, posts.content, posts.created_at, posts.id
FROM posts
JOIN users ON posts.userid = users.id
LEFT JOIN (
    SELECT postid, SUM(vote) as total_votes
    FROM post_interactions
    GROUP BY postid
) votes ON posts.id = votes.postid
ORDER BY total_votes DESC
LIMIT 21`)
	if err != nil {
		fmt.Println("error querying database")
		return
	}

	defer rows.Close()
	// if rows.Next() {
	// 	fmt.Println(models.ErrNoResultFound)
	// 	return
	// }

	arr := []*postInfo{}
	for rows.Next() {
		temp := ""
		p := postInfo{}
		err := rows.Scan(&p.PublisherName, &p.PostTitle, &p.PostBody, &p.PostDate, &temp)
		if err != nil {
			log.Printf("error reading c: %v", err)
			continue
		}

		num := noComments(temp)
		p.PostComments = num

		arr = append(arr, &p)
	}

	err = templates.LoadPage(w, "pages/home.html", info{LoggedIn: isLogged, Title: "Home", ArrInfo: arr, UsernameForNav: username, IdForNav: id})

	if err != nil {
		fmt.Printf("err: %v\n", err)
		templates.LoadPage(w, "pages/error.html", err)
		return
	}
}
