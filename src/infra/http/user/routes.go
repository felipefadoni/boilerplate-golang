package user

import "github.com/gin-gonic/gin"

func UserRoutes(r *gin.Engine) {
	userRoute := r.Group("/user")
	userRoute.POST("/create", CreateUserController)
	userRoute.DELETE("/delete/:id", DeleteUserController)
}
