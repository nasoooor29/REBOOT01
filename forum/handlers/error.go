package handlers

import (
	"db-test/templates"
	"fmt"
	"net/http"
)

func HandleErrorPage(status int, title, desc, shortDesc string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		err := templates.LoadPage(w, "pages/error.html", struct {
			Title     string
			Desc      string
			ShortDesc string
			Status    int
		}{
			Title:     title,
			Desc:      desc,
			ShortDesc: shortDesc,
			Status:    status,
		})
		fmt.Printf("err: %v\n", err)
	}
}

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	HandleErrorPage(
		http.StatusNotFound,
		"Not Found",
		"The page you are looking for not found",
		"Not Found",
	)(w, r)
}

func HandleInternalServerError(w http.ResponseWriter, r *http.Request) {
	HandleErrorPage(
		http.StatusInternalServerError,
		"Internal Server Error",
		"Something went wrong",
		"Internal Server Error",
	)(w, r)
}

func HandleBadRequest(w http.ResponseWriter, r *http.Request) {
	HandleErrorPage(
		http.StatusBadRequest,
		"Bad Request",
		"Your request is invalid",
		"Bad Request",
	)(w, r)
}