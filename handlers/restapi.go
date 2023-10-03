package handlers

import (
	movies "assignment1/connectors"
	"assignment1/structs"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartRestApi(movieDb movies.MovieDb) {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true
	router.Use(cors.New(config))
	router.GET("/movies", func(context *gin.Context) { // /movies?Offset=10&Limit=10
		arg := structs.Pagination{}
		context.ShouldBindQuery(&arg)

		res, err := movieDb.All(arg)

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
		var json structs.Movie
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
