package routes

import "github.com/gin-gonic/gin"

func AddRoutes(Route *gin.RouterGroup) {
	UserRoute(Route)
}
