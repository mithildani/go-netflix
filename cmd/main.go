package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mithildani/go-netflix/internal/routes"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	routes.RegisterMovieRoutes(v1)
	router.Run("localhost:8080")
}
