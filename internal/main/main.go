package main

import (
	"fmt"
	"net/http"
	"olx-clone-server/internal/handlers"
	"olx-clone-server/internal/models"
)

func main() {
	models.InitDB()
	http.HandleFunc("/", handlers.GetAllPostings)
	fmt.Println("Listening on port :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}