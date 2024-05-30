package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

var err error

func Init() {
	connStr := "user=postgres dbname=recipeapp password=milanpoudel host=localhost sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping()
	if err != nil {
		fmt.Println("Error while connecting to database")
		panic(err)
	}
	fmt.Println("Connected to database")
}
