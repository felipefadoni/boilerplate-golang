package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "create user",
	})
}
