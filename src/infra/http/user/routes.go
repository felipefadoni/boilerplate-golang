package user

import "github.com/gin-gonic/gin"

func UserRoutes(r *gin.Engine) {
	userRoute := r.Group("/user")
	userRoute.GET("/", GetAllUserController)
	userRoute.POST("/create", CreateUserController)
	userRoute.DELETE("/delete/:id", DeleteUserController)
}
