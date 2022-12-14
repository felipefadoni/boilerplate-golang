package user

import (
	"net/http"

	user "github.com/felipefadoni/boilerplate-golang/src/domain/user/usecases"
	validatorUser "github.com/felipefadoni/boilerplate-golang/src/infra/http/validator/user"
	"github.com/gin-gonic/gin"
)

func GetAllUserController(c *gin.Context) {
	pageQ := c.Query("page")
	limitQ := c.Query("limit")

	page, limit, err := validatorUser.ValidationGetAllUser(pageQ, limitQ)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := user.GetAllUserUseCase(page, limit)

	c.JSON(http.StatusOK, gin.H{
		"message": "Get All Users",
		"data":    result,
	})
}
