package routes

import (
	"scrape-film/handler"
	"scrape-film/repository"

	"github.com/gin-gonic/gin"
)

func MovieRoute(r *gin.RouterGroup) {

	movieRepo := repository.NewMovieRepository()
	movieHandler := handler.NewMovieHandler(*movieRepo)

	movie := r.Group("/movie")

	movie.GET("/scrape", movieHandler.StartScrape)
	movie.GET("/scrape/all", movieHandler.GetAllMovies)
}
