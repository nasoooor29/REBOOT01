package main

import (
	"fmt"
	"net/http"
)

const (
	PORT = ":8081"
)

var artists []Artist
var locations Locations
var relations Relations

func main() {
	go func() {

		//Get Artist Pointer
		artistsPTR, err := GetInfo[[]Artist](APIEndpoints["artists"])
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		artists = *artistsPTR
		//Get Loctaion Pointer
		locationPTR, err := GetInfo[Locations](APIEndpoints["locations"])
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		locations = *locationPTR

		//Get Relation Pointer
		relationPTR, err := GetInfo[Relations](APIEndpoints["relations"])
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		relations = *relationPTR
	}()

	stack := CreateStack(
		Logging,
		Recovery,
		CheckTemplates,
	)

	mux := http.NewServeMux()
	AddEndpoints(mux)
	fmt.Printf("server running in port: %v\n", PORT)
	err := http.ListenAndServe(PORT, stack(mux))
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

}

func AddEndpoints(mux *http.ServeMux) {

	mux.HandleFunc("GET /400", func(w http.ResponseWriter, r *http.Request) {
		ErrorMessagePage(w, 400, "you have a bad request message")
	})
	mux.HandleFunc("GET /404", func(w http.ResponseWriter, r *http.Request) {
		ErrorMessagePage(w, 404, "page not found message")
	})
	mux.HandleFunc("GET /panic", func(w http.ResponseWriter, r *http.Request) {
		panic("panic message here")
	})

	mux.HandleFunc("GET /", handleIndex)
	mux.HandleFunc("GET /artist/{id}", handleArtist)

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(TEMPLATE_FS)))
}
