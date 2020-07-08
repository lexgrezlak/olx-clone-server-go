package main

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"fmt"
	"olx-clone-server/config"
	"os"
)

type Config struct {
	Database string `json:"secret"`
}

func main() {
	var err error

	c := Config{}
	//if err = config.LoadFile("env.json", &c); err != nil {
	//	panic(err)
	//}

	if err = config.LoadConfig("env.json", "database", &c); err != nil {
		panic(err)
	}

	fmt.Println(c, os.Getenv("secret"))

	dbpool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	defer dbpool.Close()
	//var greeting string
	//err = dbpool.QueryRow(context.Background(), "SELECT 'hello world'").Scan(&greeting)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	//	os.Exit(1)
	//}
	//
	//fmt.Println(greeting)
}