package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"olx-clone-server/internal/handlers"
	"olx-clone-server/internal/services"
	"time"
)

func main() {
	services.InitDB()
	r := mux.NewRouter()
	r.HandleFunc("/postings", handlers.Postings)
	//r.Use(middleware.Logger)
	r.Use(mux.CORSMethodMiddleware(r))
	srv := &http.Server{
		Handler: r,
		Addr: "127.0.0.1:4000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	fmt.Println("Listening at:", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}