package server

import (
	"fmt"
	"os"
	databases "pp-model-shop-backend/pkg/database"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Instance *gorm.DB
)

func CreateDataBase() {
	godotenv.Load("../.env")
	dsn := os.Getenv("DATABASE")
	fmt.Printf("%v", &dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&databases.User{}, &databases.Address{}, &databases.Address{}, &databases.ImageProduct{},
		&databases.Cart{}, &databases.Order{}, &databases.OrderItem{})
	Instance = db

}
