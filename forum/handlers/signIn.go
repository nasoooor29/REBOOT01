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

func HandleSignInPageForm(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	if strings.TrimSpace(email) == "" || strings.TrimSpace(password) == "" {
		err := templates.LoadPage(w, "pages/signin.html", struct {
			Msg   string
			Title string
		}{
			Msg:   "All fields must be filled!",
			Title: "signin",
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
		if user.Email == email {
			match := utils.CheckPasswordHash(password, user.Password)
			if match {
				err := utils.GenerateCookie(w, user.ID)
				if err != nil {
					fmt.Println(err)
				}
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			} else {
				err := templates.LoadPage(w, "pages/signin.html", struct {
					Msg   string
					Title string
				}{
					Msg:   "wrong password!",
					Title: "signin",
				})
				if err != nil {
					fmt.Printf("err: %v\n", err)
					HandleBadRequest(w, r)
					return
				}
				return
			}
		}
	}
	err = templates.LoadPage(w, "pages/signin.html", struct {
		Msg   string
		Title string
	}{
		Msg:   "email is not found in our database!",
		Title: "signin",
	})
	if err != nil {
		HandleBadRequest(w, r)
		return
	}
	return
}

func HandleSignInPage(w http.ResponseWriter, r *http.Request) {
	err := templates.LoadPage(w, "pages/signin.html", nil)
	if err != nil {
		HandleInternalServerError(w, r)
		return
	}

}
