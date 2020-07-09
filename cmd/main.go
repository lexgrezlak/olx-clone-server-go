package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"olx-clone-server/internal/handlers"
	//"olx-clone-server/internal/middleware"
	"olx-clone-server/internal/models"
	"time"
)

func main() {
	models.InitDB()
	r := mux.NewRouter()
	r.HandleFunc("/postings", handlers.Postings)

	//r.Use(middleware.Logger)
	//h := middleware.ApplyMiddleware(
	//	handlers.GetAllPostings,
	//	middleware.Logger(log.New(os.Stdout, "", 0)),
	//	middleware.SetID(1),
	//	)
	//http.HandleFunc("/", h)
	//fmt.Println("Listening on port :3333")
	//err := http.ListenAndServe(":3333",  nil)
	//panic(err)
	srv := &http.Server{
		Handler: r,
		Addr: "127.0.0.1:4000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	fmt.Println("Listening at:", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}