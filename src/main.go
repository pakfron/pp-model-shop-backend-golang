package main

import (
	"net/http"
	"pp-model-shop-backend/auth-controller"
	pp_model_schema "pp-model-shop-backend/database"

	"github.com/gin-gonic/gin"
)

func main() {
	pp_model_schema.CreateDataBase()
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/register", auth.Register)
	r.Run("localhost:8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
