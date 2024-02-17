package routes

import "github.com/gin-gonic/gin"

func Routes(Route *gin.RouterGroup) {
	UserRoute(Route)
}
