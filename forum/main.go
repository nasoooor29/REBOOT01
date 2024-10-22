package main

import (
	"db-test/db"
	"db-test/handlers"
	"db-test/middleware"
	"db-test/models"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// os.Remove(models.DB_NAME)
	// fmt.Println("DB initialized")
	defer CloseDB()
	err := db.Init()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	// db.GenerateDummyData()
	_, err = db.ReadAllCategory()
	if err == models.ErrNoResultFound {
		db.GenerateCatagories()
	}
	
	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.Auth,
		middleware.Recovery,
	)
	mux := http.NewServeMux()
	handlers.AddHandlers(mux)
	fmt.Printf("Listening on %v\n", models.DEFAULT_PORT)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	err = http.ListenAndServe(models.DEFAULT_PORT, stack(mux))
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
}

func CloseDB() {
	err := db.Database.Close()
	if err != nil {
		fmt.Printf("error closing the DB: %v\n", err)
		return
	}
}
