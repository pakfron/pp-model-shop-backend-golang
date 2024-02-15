package main

import (
	authRegister "pp-model-shop-backend/controller/auth-controller"
	pp_model_schema "pp-model-shop-backend/database"

	"github.com/gin-gonic/gin"
)

func main() {
	pp_model_schema.CreateDataBase()
	r := gin.Default()

	r.POST("/register", authRegister.Register)
	r.Run("localhost:8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
