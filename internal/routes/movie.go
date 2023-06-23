package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mithildani/go-netflix/internal/controllers"
)

func RegisterMovieRoutes(rg *gin.RouterGroup) {
	movie := rg.Group("/movie")
	// movie.POST("/movie", addMovie)
	// movie.GET("/", getAllMovies)
	movie.GET("/:movieid", controllers.GetMovie)
	// movie.PUT("/movie", updateMovie)
	// movie.DELETE("/movie", deleteMovie)
}
