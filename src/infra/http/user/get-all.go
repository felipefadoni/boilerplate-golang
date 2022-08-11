package user

import (
	"net/http"

	"github.com/felipefadoni/boilerplate-golang/src/modules/user"
	"github.com/gin-gonic/gin"
)

func GetAllUserController(c *gin.Context) {
	users := user.GetAllUserModule()

	c.JSON(http.StatusOK, gin.H{
		"message": "Get All Users",
		"data":    users,
	})
}
