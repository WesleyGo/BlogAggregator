package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/WESLEYGO/BlogAggregator/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	godotenv.Load()
	port := os.Getenv("PORT")
	fmt.Println("Port is: ", port)
	fmt.Println("Hello, World!")

	db, err := sql.Open("postgres", os.Getenv("CONN"))

	if err != nil {
		fmt.Println("Error connecting to the database: ", err)
	} else {
		dbQueries := database.New(db)
		apiCfg := apiConfig{
			DB: dbQueries,
		}
		fmt.Println("Connected to the database")
		InitServer(port, &apiCfg)
	}
}
