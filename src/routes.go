package src

import (
	"net/http"
	"os"
	"time"

	"github.com/felipefadoni/boilerplate-golang/src/infra/http/user"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		currentTime := time.Now()
		apiVersion := os.Getenv("API_VERSION")
		c.JSON(http.StatusOK, gin.H{
			"message": "API STARTED",
			"version": apiVersion,
			"date":    currentTime,
		})
	}).Static("/static", "./static")

	user.UserRoutes(r)

	return r
}
