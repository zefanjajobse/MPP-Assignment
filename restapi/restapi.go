package restapi

import (
	"assignment1/movies"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start(movieDb movies.MovieDb) {
	router := gin.Default()
	router.GET("/movies", func(context *gin.Context) {
		res, err := movieDb.All()

		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, res)
	})
	router.GET("/movies/:id", func(context *gin.Context) {
		res, err := movieDb.FindOne(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, res)
	})
	router.POST("/movies", func(context *gin.Context) {
		var json movies.Movie
		if err := context.ShouldBindJSON(&json); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		movie, err := movieDb.Insert(json)

		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, movie)
	})
	router.DELETE("/movies/:id", func(context *gin.Context) {
		movie_id := context.Param("id")
		movieDb.Delete(movie_id)
		context.JSON(http.StatusNoContent, gin.H{"deleted": movie_id})
	})

	router.Run("localhost:8090")
}
