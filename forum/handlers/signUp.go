package handlers

import (
	"db-test/db"
	"db-test/models"
	"db-test/templates"
	"db-test/utils"
	"fmt"
	"net/http"
	"strings"
)

func HandleSignUpForm(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm password")

	if strings.TrimSpace(email) == "" || strings.TrimSpace(username) == "" ||strings.TrimSpace(password) == "" {
		err := templates.LoadPage(w, "pages/signup.html", struct {
			Msg   string
			Title string
		}{
			Msg:   "All fields must be filled!",
			Title: "signup",
		})
		if err != nil {
			fmt.Printf("err1: %v\n", err)
			HandleBadRequest(w, r)
			return
		}
		return
	}

	if strings.TrimSpace(password) != strings.TrimSpace(confirmPassword) {
		err := templates.LoadPage(w, "pages/signup.html", struct {
			Msg   string
			Title string
		}{
			Msg:   "Password does not match!",
			Title: "signup",
		})
		if err != nil {
			HandleBadRequest(w, r)
			return
		}
		return
	}

	if strings.Contains(password, " ") {
		err := templates.LoadPage(w, "pages/signup.html", struct {
			Msg   string
			Title string
		}{
			Msg: "passward must not contains spaces!",
			Title: "signup",
		})
		if err != nil {
			HandleBadRequest(w, r)
			return
		}
		return
	}

	users, err := db.ReadAllUser()
	if err != nil && err != models.ErrNoResultFound {
		HandleInternalServerError(w, r)
	}

	for _, user := range users {
		if user.Username == username || user.Email == email {
			err := templates.LoadPage(w, "pages/signup.html", struct {
				Msg   string
				Title string
			}{
				Msg:   "username or email is already exits!",
				Title: "signup",
			})
			if err != nil {
				HandleBadRequest(w, r)
				return
			}
			return
		}
	}

	if len(password) < 8 {
		err := templates.LoadPage(w, "pages/signup.html", struct {
			Msg   string
			Title string
		}{
			Msg:   "Password must be at least 8 characters",
			Title: "signup",
		})
		if err != nil {
			HandleBadRequest(w, r)
			return
		}
		return
	}

	if len(username) > 15 {
		err := templates.LoadPage(w, "pages/signup.html", struct {
			Msg   string
			Title string
		}{
			Msg:   "username must not exceeds 15 characters",
			Title: "signup",
		})
		if err != nil {
			HandleBadRequest(w, r)
			return
		}
		return
	}

	if strings.Contains(username, " ") {
		err := templates.LoadPage(w, "pages/signup.html", struct {
			Msg   string
			Title string
		}{
			Msg:   "username must not has spaces",
			Title: "signup",
		})
		if err != nil {
			HandleBadRequest(w, r)
			return
		}
		return
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		fmt.Printf("err2: %v\n", err)
		templates.LoadPage(w, "error.html", err)
		return
	}

	id, err := db.CreateUser(username, email, string(hashedPassword))
	if err != nil {
		fmt.Printf("err3: %v\n", err)
		templates.LoadPage(w, "error.html", err)
		return
	}
	utils.GenerateCookie(w, id)
	fmt.Println("The cookie has been set sucessfully")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func HandleSignUpPage(w http.ResponseWriter, r *http.Request) {
	err := templates.LoadPage(w, "pages/signup.html", nil)
	if err != nil {
		HandleInternalServerError(w, r)
		return
	}
}
