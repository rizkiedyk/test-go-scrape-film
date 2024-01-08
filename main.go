package main

import (
	"log"
	"scrape-film/connection"
	"scrape-film/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connection.DatabaseConnect()

	r := gin.Default()

	routes.MovieRoute(&r.RouterGroup)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
