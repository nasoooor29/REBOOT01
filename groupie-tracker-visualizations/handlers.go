package main

import (
	// "fmt"

	"fmt"
	"html/template"
	"net/http"
	"strconv"
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

	err = tmpl.Execute(w, artists)
	if err != nil {
		ErrorMessagePage(w, http.StatusInternalServerError, "Error executing template")
		return
	}
}

func handleArtist(w http.ResponseWriter, r *http.Request) {
	//Validate ID
	artistID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		ErrorMessagePage(w, http.StatusNotFound, "Id not present")
		return
	}

	artist, err := filterArtist(&artists, func(artist Artist) bool { return artist.ID == artistID })
	if err != nil {
		ErrorMessagePage(w, http.StatusNotFound, "id not found")
		return
	}

	relation, err := filterRelation(&relations, func(relation Relation) bool { return relation.ID == artistID })
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	info := struct {
		ArtistObj   *Artist
		RelationObj map[string][]string
	}{
		ArtistObj:   artist,
		RelationObj: relation.DatesLocations,
	}

	tmpl, err := template.ParseFiles(ArtistTemplate)
	if err != nil {

		ErrorMessagePage(w, http.StatusInternalServerError, "Error parsing template")
		return
	}

	err = tmpl.Execute(w, info)

	if err != nil {

		ErrorMessagePage(w, http.StatusInternalServerError, "Error executing template")
		return
	}
}
