package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"olx-clone-server/internal/handler"
	"olx-clone-server/internal/middleware"
	"olx-clone-server/internal/service"
	"time"
	"github.com/rs/cors"
)

func main() {
	service.InitDB()
	r := mux.NewRouter()
	r.HandleFunc("/postings", handler.Postings).Methods("GET", "POST")
	//r.Use(middleware.Logger)

	r.Use(mux.CORSMethodMiddleware(r))


	// For dev only - Set up CORS so our client can consume the API
	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PATCH", "PUT"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	srv := &http.Server{
		Handler: corsWrapper.Handler(r),
		Addr: "127.0.0.1:4000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	fmt.Println("Listening at:", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}