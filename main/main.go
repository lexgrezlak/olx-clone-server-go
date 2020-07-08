package main

import (
	"context"
	"fmt"
	"olx-clone-server/database"
	"os"
)


func main() {
	var err error

	c := database.Config{}
	ctx := context.Background()
	if err = database.LoadConfig("env.json", "database", &c); err != nil {
		panic(err)
	}

	db, err := database.NewFromEnv(ctx, &c)
	if err != nil {
		panic(err)
	}
	defer db.Close(ctx)

	var email string
	err = db.Pool.QueryRow(ctx, "SELECT email FROM user").Scan(&email);
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	fmt.Println(email)


	//dbpool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	//}
	//defer dbpool.Close()
	//var greeting string
	//err = dbpool.QueryRow(context.Background(), "SELECT 'hello world'").Scan(&greeting)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	//	os.Exit(1)
	//}
	//
	//fmt.Println(greeting)
}