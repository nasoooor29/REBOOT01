package handlers

import (
	"db-test/db"
	"db-test/models"
	"db-test/utils"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func HandleCreatePostForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	categoriesUserInput := r.Form["Category"]
	title := r.FormValue("post-title")
	body := r.FormValue("body")

	categories := []models.Category{}
	for _, categoryID := range categoriesUserInput {
		id, err := strconv.Atoi(categoryID)
		if err != nil {
			fmt.Print(err)
			HandleInternalServerError(w, r)
			return
		}
		category, err := db.ReadCategory(id)
		if err != nil {
			//	fmt.Println("heyyy")
			fmt.Print(err)
			HandleInternalServerError(w, r)
			return
		}
		categories = append(categories, *category)
	}

	cookie, err := utils.CheckIfAuth(r)
	if err != nil {
		//	fmt.Println("VAT U DO??")
		return
	}

	post := models.Post{
		Title:      title,
		Content:    body,
		Categories: categories,
		UserID:     cookie.UserID,
	}

	if strings.TrimSpace(post.Title) == "" || strings.TrimSpace(post.Content) == "" {
		HandleBadRequest(w, r)
	}

	err = db.CreatePost(post)
	if err != nil {
		fmt.Println("Error creating post:", err)
		HandleInternalServerError(w, r)
		return
	}

	// Get the last created post
	var postID int
	err = db.Database.QueryRow("SELECT id FROM posts WHERE title = ? AND content = ? AND userid = ? ORDER BY created_at DESC LIMIT 1",
		post.Title, post.Content, post.UserID).Scan(&postID)
	if err != nil {
		fmt.Println("Error retrieving post ID:", err)
		HandleInternalServerError(w, r)
		return
	}

	// Redirect the user to its created post page
	http.Redirect(w, r, fmt.Sprintf("/post/%d", postID), http.StatusSeeOther)
}
