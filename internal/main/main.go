package main

import (
	"fmt"
	"net/http"
	"olx-clone-server/internal/database"
	"olx-clone-server/internal/handlers"
	"olx-clone-server/internal/models"
)

var db *database.DB

func main() {
	models.InitDB()
	http.HandleFunc("/", handlers.GetAllPostings)
	fmt.Println("Listening on port :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)

	//var title string
	//err := db.Pool.QueryRow(ctx, "SELECT title FROM posting").Scan(&title);
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	//}

}