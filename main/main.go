package main

import (
	"fmt"
	"olx-clone-server/database"
)


func main() {
	var err error

	c := database.Config{}
	if err = database.LoadConfig("env.json", "database", &c); err != nil {
		panic(err)
	}

	fmt.Println(c)

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