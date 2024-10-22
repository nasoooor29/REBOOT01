package main

import (
	ascii "a1/ascii-art"
	"fmt"
	"html/template"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorMessagePage(w, http.StatusNotFound, "the page you are looking for is not here")
		return
	}
	tmpl, err := template.ParseFiles(IMPORTANT_FILES[http.StatusOK])
	if err != nil {

		ErrorMessagePage(w, http.StatusInternalServerError, "Error parsing template")
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		ErrorMessagePage(w, http.StatusInternalServerError, "Error executing template")
		return
	}
}

func handleASCII(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	inputString := r.FormValue("InputValue")
	inputFont := r.FormValue("InputFont")

	Output, err := ascii.Output(inputString, inputFont)
	if err != nil {
		ErrorMessagePage(w, http.StatusNotAcceptable, "Only ASCII Characters accepted")
		fmt.Println(err.Error())
		return

	}
	// fmt.Println(Output)

	data := struct {
		Value string
	}{
		Value: Output,
	}
	tmpl, err := template.ParseFiles(IMPORTANT_FILES[http.StatusOK])
	if err != nil {

		ErrorMessagePage(w, http.StatusInternalServerError, "Error parsing template")
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		ErrorMessagePage(w, http.StatusInternalServerError, "Error executing template")
		return
	}

}
