package services

import (
	"context"
	"olx-clone-server/internal/database"
)

var db *database.DB

// Inits the database with config from the root folder
func InitDB() {
	var err error
	c := database.Config{}
	if err = database.LoadConfig("env.json", "database", &c); err != nil {
		panic(err)
	}

	db, err = database.NewFromEnv(context.Background(), &c)
	if err != nil {
		panic(err)
	}
}

