package handlers

import (
	"db-test/db"
	"db-test/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func HandleLikeComment(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Fatima")
	isLogged := true
	cookie, err := utils.CheckIfAuth(r)
	if err != nil {
		//fmt.Println("Fatima Kitchen")
		isLogged = false
		http.Redirect(w, r, "/signIn", http.StatusSeeOther)
	}

	commentNumber, err := strconv.Atoi(r.PathValue("CommentID"))
	if err != nil {
		HandleBadRequest(w, r)
		return
	}

	var likeCheck int
	var dislikeCheck int
	db.Database.QueryRow(`SELECT COUNT(*) FROM comment_interactions WHERE commentid = ? AND userid = ? AND interaction = 0`, commentNumber, cookie.UserID).Scan(&likeCheck)
	fmt.Println(likeCheck)
	db.Database.QueryRow(`SELECT COUNT(*) FROM comment_interactions WHERE commentid = ? AND userid = ? AND interaction = 1`, commentNumber, cookie.UserID).Scan(&dislikeCheck)
	if dislikeCheck != 0 {
		db.Database.Exec(`DELETE FROM comment_interactions WHERE userid = ? AND commentid = ?`, cookie.UserID, commentNumber)
	}
	if likeCheck == 1 {
		db.Database.Exec(`DELETE FROM comment_interactions WHERE userid = ? AND commentid = ?`, cookie.UserID, commentNumber)
	} else {
		_, err := db.CreateCommentInteraction(0, cookie.UserID, commentNumber)
		if err != nil {
			HandleBadRequest(w, r)
			return
		}
		//	fmt.Println("LIKE ADDED")
	}

	fmt.Println(likeCheck)
	fmt.Println(dislikeCheck)

	fmt.Println(cookie.UserID, "AND", commentNumber)
	fmt.Println(isLogged)

	var totalLikes int
	var totalDislikes int
	db.Database.QueryRow(`SELECT COUNT(*) FROM comment_interactions WHERE commentid = ? AND interaction = 0`, commentNumber).Scan(&totalLikes)
	db.Database.QueryRow(`SELECT COUNT(*) FROM comment_interactions WHERE commentid = ? AND interaction = 1`, commentNumber).Scan(&totalDislikes)

	data := likesData{
		Likes:    totalLikes,
		Dislikes: totalDislikes,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func HandleDislikeComment(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println("Fatima")
	isLogged := true
	cookie, err := utils.CheckIfAuth(r)
	if err != nil {
		//fmt.Println("Fatima Kitchen")
		isLogged = false
		http.Redirect(w, r, "/signIn", http.StatusSeeOther)
	}

	commentNumber, err := strconv.Atoi(r.PathValue("CommentID"))
	if err != nil {
		HandleBadRequest(w, r)
		fmt.Println(err)
		return
	}

	var likeCheck int
	var dislikeCheck int
	db.Database.QueryRow(`SELECT COUNT(*) FROM comment_interactions WHERE commentid = ? AND userid = ? AND interaction = 0`, commentNumber, cookie.UserID).Scan(&likeCheck)
	fmt.Println(likeCheck)
	db.Database.QueryRow(`SELECT COUNT(*) FROM comment_interactions WHERE commentid = ? AND userid = ? AND interaction = 1`, commentNumber, cookie.UserID).Scan(&dislikeCheck)
	if likeCheck != 0 {
		db.Database.Exec(`DELETE FROM comment_interactions WHERE userid = ? AND commentid = ?`, cookie.UserID, commentNumber)
	}
	if dislikeCheck == 1 {
		db.Database.Exec(`DELETE FROM comment_interactions WHERE userid = ? AND commentid = ?`, cookie.UserID, commentNumber)
	} else {
		_, err := db.CreateCommentInteraction(1, cookie.UserID, commentNumber)
		if err != nil {
			HandleBadRequest(w, r)
			return
		}
		fmt.Println("LIKE ADDED")
	}

	fmt.Println(likeCheck)
	fmt.Println(dislikeCheck)

	fmt.Println(cookie.UserID, "AND", commentNumber)
	fmt.Println(isLogged)

	var totalLikes int
	var totalDislikes int
	db.Database.QueryRow(`SELECT COUNT(*) FROM comment_interactions WHERE commentid = ? AND interaction = 0`, commentNumber).Scan(&totalLikes)
	db.Database.QueryRow(`SELECT COUNT(*) FROM comment_interactions WHERE commentid = ? AND interaction = 1`, commentNumber).Scan(&totalDislikes)

	data := likesData{
		Likes:    totalLikes,
		Dislikes: totalDislikes,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
