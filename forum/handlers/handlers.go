package handlers

import (
	"db-test/static"
	"net/http"
)

func AddHandlers(mux *http.ServeMux) {
	mux.HandleFunc("GET /", HandleHomePage)
	mux.HandleFunc("GET /signIn", HandleSignInPage)
	mux.HandleFunc("POST /signIn", HandleSignInPageForm)
	mux.HandleFunc("GET /signUp", HandleSignUpPage)
	mux.HandleFunc("POST /signUp", HandleSignUpForm)
	mux.HandleFunc("GET /logout", HandleLogOut)
	mux.HandleFunc("GET /posts/{pgNumber}", HandlePostsPage)
	mux.HandleFunc("GET /post/{id}", HandlePostPage)
	mux.HandleFunc("POST /post/", HandlePostPageForm)
	mux.HandleFunc("POST /like/{PostID}", HandleLikePost)
	mux.HandleFunc("POST /dislike/{PostID}", HandleDislikePost)
	mux.HandleFunc("POST /likeComment/{CommentID}", HandleLikeComment)
	mux.HandleFunc("POST /dislikeComment/{CommentID}", HandleDislikeComment)
	mux.HandleFunc("GET /profile/{id}", HandleProfilePage)
	mux.HandleFunc("POST /createPost", HandleCreatePostForm)
	// db.Database.Query()
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServerFS(static.Assets)))
}
