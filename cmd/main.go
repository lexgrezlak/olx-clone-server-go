package main

import (
	"fmt"
	"log"
	"net/http"
	"olx-clone-server/internal/handlers"
	"olx-clone-server/internal/middleware"
	"olx-clone-server/internal/models"
	"os"
)

func main() {
	models.InitDB()
	h := middleware.ApplyMiddleware(
		handlers.GetAllPostings,
		middleware.Logger(log.New(os.Stdout, "", 0)),
		middleware.SetID(1),
		)
	http.HandleFunc("/", h)
	fmt.Println("Listening on port :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}