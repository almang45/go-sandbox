package main

import (
	"github.com/almang45/go-sandbox/amdb/services"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("api/v1-go")
	{
		v1.GET("/list-movies", services.GetMovies)
		v1.GET("/new-r-movies", services.GetNewReleaseMovies)
		v1.GET("/new-d-movies", services.GetNewDownloadMovies)
		v1.GET("/movie/:id", services.GetMovieByImdbId)
		v1.GET("/movie-a/:actor", services.GetMoviesByActor)
		v1.GET("/movie-d/:director", services.GetMoviesByDirector)
		v1.GET("/movie-g/:genre", services.GetMoviesByGenre)
		v1.POST("/movie", services.SaveOrUpdateMovies)
		v1.POST("/movie-p", services.UpdateMoviePoster)
		//---
		v1.POST("/login", services.Login)
		v1.POST("/user", services.CreateUser)
		v1.GET("/user/:email", services.GetUserByEmail)
		//---
		// v1.GET("/image", services.ImageConverter)
	}

	r.Run(":8080")
}
