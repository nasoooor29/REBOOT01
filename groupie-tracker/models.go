package main

import "net/http"

var IMPORTANT_FILES = map[int]string{
	400: "templates/400.tmpl.html",
	404: "templates/404.tmpl.html",
	500: "templates/500.tmpl.html",
	200: "templates/index.html",
}

var ArtistTemplate = "templates/artist.html"

var APIEndpoints = map[string]string{
	"artists":   "https://groupietrackers.herokuapp.com/api/artists",
	"locations": "https://groupietrackers.herokuapp.com/api/locations",
	"dates":     "https://groupietrackers.herokuapp.com/api/dates",
	"relations": "https://groupietrackers.herokuapp.com/api/relation",
}

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Dates struct {
	Index []Date
}
type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Locations struct {
	Index []Location
}
type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Relations struct {
	Index []Relation
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

var TEMPLATE_FS = http.Dir("static")
