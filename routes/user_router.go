package routes

import (
	userControllers "pp-model-shop-backend/modules/users/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.RouterGroup) {

	userRouter := router.Group("/user")
	{
		userRouter.POST("/register", userControllers.Register)
	}

}
