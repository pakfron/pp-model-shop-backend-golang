package main

import (
	"pp-model-shop-backend/modules/server"
	"pp-model-shop-backend/modules/users/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server.CreateDataBase()
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/register", controllers.Register)
	r.Run("localhost:8008") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
