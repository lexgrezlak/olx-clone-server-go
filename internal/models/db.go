package models

import (
	"context"
	"olx-clone-server/internal/database"
)

var db *database.DB

func InitDB() {
	var err error
	ctx := context.Background()
	c := database.Config{}
	if err = database.LoadConfig("../env.json", "database", &c); err != nil {
		panic(err)
	}

	db, err = database.NewFromEnv(ctx, &c)
	if err != nil {
		panic(err)
	}
}