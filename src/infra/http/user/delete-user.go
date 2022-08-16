package user

import (
	"net/http"

	validatorUser "github.com/felipefadoni/boilerplate-golang/src/infra/http/validator/user"
	"github.com/felipefadoni/boilerplate-golang/src/useCases/user"
	"github.com/gin-gonic/gin"
)

func DeleteUserController(c *gin.Context) {

	id := c.Param("id")

	err := validatorUser.ValidationDeleteUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.DeleteUserUseCase(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error deleting user"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
