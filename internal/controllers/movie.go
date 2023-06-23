package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMovie(c *gin.Context) {
	movieid := c.Param("movieid")
	response := fmt.Sprintf("This is Movie ID: %v", movieid)
	c.IndentedJSON(http.StatusOK, response)
}
