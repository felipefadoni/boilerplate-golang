package user

import (
	"net/http"

	"github.com/felipefadoni/boilerplate-golang/src/modules/user"
	"github.com/gin-gonic/gin"
)

func DeleteUserController(c *gin.Context) {

	id := c.Param("id")

	user.DeleteUserService(id)

	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted",
		"id":      id,
	})
}
