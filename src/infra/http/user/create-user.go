package user

import (
	"net/http"

	user "github.com/felipefadoni/boilerplate-golang/src/domain/user/usecases"
	"github.com/felipefadoni/boilerplate-golang/src/dto"
	validatorUser "github.com/felipefadoni/boilerplate-golang/src/infra/http/validator/user"
	"github.com/gin-gonic/gin"
)

func CreateUserController(c *gin.Context) {

	var userDTO dto.CreateUserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := validatorUser.ValidationCreateUser(userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := user.CreateUserUseCase(userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "create user",
		"data":    user,
	})
}
