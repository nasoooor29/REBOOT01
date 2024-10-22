package handlers

import (
	"db-test/db"
	"db-test/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type likesData struct {
	Likes    int `json:"likes"`
	Dislikes int `json:"dislikes"`
}

func HandleLikePost(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println("Fatima")
	isLogged := true
	cookie, err := utils.CheckIfAuth(r)
	if err != nil {
		//	fmt.Println("Fatima Kitchen")
		isLogged = false
		http.Redirect(w, r, "/signIn", http.StatusSeeOther)
	}

	postNumber, err := strconv.Atoi(r.PathValue("PostID"))
	if err != nil {
		HandleBadRequest(w, r)
		return
	}

	var likeCheck int
	var dislikeCheck int
	db.Database.QueryRow(`SELECT COUNT(*) FROM post_interactions WHERE postid = ? AND userid = ? AND vote = 0`, postNumber, cookie.UserID).Scan(&likeCheck)
	fmt.Println(likeCheck)
	db.Database.QueryRow(`SELECT COUNT(*) FROM post_interactions WHERE postid = ? AND userid = ? AND vote = 1`, postNumber, cookie.UserID).Scan(&dislikeCheck)
	if dislikeCheck != 0 {
		db.Database.Exec(`DELETE FROM post_interactions WHERE userid = ? AND postid = ?`, cookie.UserID, postNumber)
	}
	if likeCheck == 1 {
		db.Database.Exec(`DELETE FROM post_interactions WHERE userid = ? AND postid = ?`, cookie.UserID, postNumber)
	} else {
		err := db.CreatePostInteraction(0, cookie.UserID, postNumber)
		if err != nil {
			HandleBadRequest(w, r)
			return
		}
		fmt.Println("LIKE ADDED")
	}
	fmt.Println(cookie.UserID, "AND", postNumber)
	fmt.Println(isLogged)

	var totalLikes int
	var totalDislikes int
	db.Database.QueryRow(`SELECT COUNT(*) FROM post_interactions WHERE postid = ? AND vote = 0`, postNumber).Scan(&totalLikes)
	db.Database.QueryRow(`SELECT COUNT(*) FROM post_interactions WHERE postid = ? AND vote = 1`, postNumber).Scan(&totalDislikes)

	data := likesData{
		Likes:    totalLikes,
		Dislikes: totalDislikes,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func HandleDislikePost(w http.ResponseWriter, r *http.Request) {
	isLogged := true
	cookie, err := utils.CheckIfAuth(r)
	if err != nil {
		isLogged = false
		http.Redirect(w, r, "/signIn", http.StatusSeeOther)
	}

	postNumber, err := strconv.Atoi(r.PathValue("PostID"))
	if err != nil {
		HandleBadRequest(w, r)
		return
	}
	var likeCheck int
	db.Database.QueryRow(`SELECT COUNT(*) FROM post_interactions WHERE postid = ? AND userid = ? AND vote = 0`, postNumber, cookie.UserID).Scan(&likeCheck)
	var dislikeCheck int
	db.Database.QueryRow(`SELECT COUNT(*) FROM post_interactions WHERE postid = ? AND userid = ? AND vote = 1`, postNumber, cookie.UserID).Scan(&dislikeCheck)
	fmt.Println(dislikeCheck)

	if likeCheck != 0 {
		db.Database.Exec(`DELETE FROM post_interactions WHERE userid = ? AND postid = ?`, cookie.UserID, postNumber)
	}
	if dislikeCheck == 1 {
		db.Database.Exec(`DELETE FROM post_interactions WHERE userid = ? AND postid = ?`, cookie.UserID, postNumber)
	} else {
		err := db.CreatePostInteraction(1, cookie.UserID, postNumber)
		if err != nil {
			HandleBadRequest(w, r)
			return
		}
		fmt.Println("DISLIKE ADDED")
	}
	fmt.Println(cookie.UserID, "AND", postNumber)
	fmt.Println(isLogged)

	var totalLikes int
	var totalDislikes int
	db.Database.QueryRow(`SELECT COUNT(*) FROM post_interactions WHERE postid = ? AND vote = 0`, postNumber).Scan(&totalLikes)
	db.Database.QueryRow(`SELECT COUNT(*) FROM post_interactions WHERE postid = ? AND vote = 1`, postNumber).Scan(&totalDislikes)

	data := likesData{
		Likes:    totalLikes,
		Dislikes: totalDislikes,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
