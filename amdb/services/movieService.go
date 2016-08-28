package services

import (
	"github.com/almang45/go-sandbox/amdb/models"
	"github.com/almang45/go-sandbox/amdb/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var movieDbmap = utils.InitMovieDb()

func GetMovieByImdbId(c *gin.Context) {
	imdbId := c.Params.ByName("id")
	movie := getMovieByImdbId(imdbId)
	if (models.MovieDetails{}) != movie {
		c.JSON(200, movie)
	} else {
		c.JSON(404, gin.H{"error": "Movie not found."})
	}
}

func getMovieByImdbId(imdbId string) models.MovieDetails {
	var movie models.MovieDetails
	movieDbmap.SelectOne(&movie, "SELECT * FROM movies WHERE imdbId=$1", imdbId)
	return movie
}

func getMovieByTitleAndYear(title string, year string) models.MovieDetails {
	var movie models.MovieDetails
	movieDbmap.SelectOne(&movie, "SELECT * FROM movies WHERE title=$1 AND year=$1", title, year)
	return movie
}

func GetMovies(c *gin.Context) {
	var movies []models.MovieDetails
	_, err := movieDbmap.Select(&movies, "SELECT * FROM movies ORDER BY type, title")
	if err == nil {
		c.JSON(200, movies)
	} else {
		c.JSON(404, gin.H{"error": "No movies recorded."})
	}
}

func GetMoviesByActor(c *gin.Context) {
	actor := "%" + c.Params.ByName("actor") + "%"
	var movies []models.MovieDetails
	_, err := movieDbmap.Select(&movies, "SELECT * FROM movies WHERE LOWER(actors) LIKE $1 ORDER BY type, title", actor)
	if err == nil {
		c.JSON(200, movies)
	} else {
		c.JSON(404, gin.H{"error": "No movies recorded with that actor."})
	}
}

func GetMoviesByDirector(c *gin.Context) {
	director := "%" + c.Params.ByName("director") + "%"
	var movies []models.MovieDetails
	_, err := movieDbmap.Select(&movies, "SELECT * FROM movies WHERE LOWER(director) LIKE $1 ORDER BY type, title", director)
	if err == nil {
		c.JSON(200, movies)
	} else {
		c.JSON(404, gin.H{"error": "No movies recorded with that Director."})
	}
}

func GetMoviesByGenre(c *gin.Context) {
	genre := "%" + c.Params.ByName("genre") + "%"
	var movies []models.MovieDetails
	_, err := movieDbmap.Select(&movies, "SELECT * FROM movies WHERE LOWER(genre) LIKE $1 ORDER BY type, title", genre)
	if err == nil {
		c.JSON(200, movies)
	} else {
		c.JSON(404, gin.H{"error": "No movies recorded with that Genre."})
	}
}

func GetNewDownloadMovies(c *gin.Context) {
	var movies []models.MovieDetails
	_, err := movieDbmap.Select(&movies, "SELECT * FROM movies WHERE is_new_download=true ORDER BY type, title")
	if err == nil {
		c.JSON(200, movies)
	} else {
		c.JSON(404, gin.H{"error": "No movies recorded."})
	}
}

func GetNewReleaseMovies(c *gin.Context) {
	var movies []models.MovieDetails
	_, err := movieDbmap.Select(&movies, "SELECT * FROM movies WHERE is_new_release=true ORDER BY type, title")
	if err == nil {
		c.JSON(200, movies)
	} else {
		c.JSON(404, gin.H{"error": "No movies recorded."})
	}
}

func SaveOrUpdateMovies(c *gin.Context) {
	var movies []models.MovieDetails
	c.Bind(&movies)
	for _, movie := range movies {
		movieByImdbId := getMovieByImdbId(movie.ImdbId)
		if (models.MovieDetails{}) != movieByImdbId {
			updateMovie(movie, movieByImdbId.Id)
		} else {
			movieByTitleYear := getMovieByTitleAndYear(movie.Title, movie.Year)
			if (models.MovieDetails{}) != movieByTitleYear {
				updateMovie(movie, movieByTitleYear.Id)
			} else {
				err := movieDbmap.Insert(&movie)
				utils.CheckError(err, "Insert failed!")
			}
		}
	}
}

func updateMovie(movie models.MovieDetails, id int) {
	movie.Id = id
	_, err := movieDbmap.Update(&movie)
	utils.CheckError(err, "Update failed!")
}

func UpdateMoviePoster(c *gin.Context) {
	var movies []models.MovieDetails
	c.Bind(&movies)
	for _, movie := range movies {
		oldMovie := getMovieByImdbId(movie.ImdbId)
		if (models.MovieDetails{}) != oldMovie {
			oldMovie.Poster = movie.Poster
			updateMovie(oldMovie, oldMovie.Id)
		}
	}
}
