package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return r
}
