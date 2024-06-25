package main

import (
	"fmt"
	"os"

	"github.com/WESLEYGO/BlogAggregator/internal/http"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	port := os.Getenv("PORT")
	fmt.Println("Port is: ", port)
	fmt.Println("Hello, World!")
	http.InitServer(port)
}
