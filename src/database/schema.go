package pp_model_schema

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type RoleType string

const (
	ADMIN = "ADMIN"
	USER  = "USER"
)

type User struct {
	gorm.Model
	UserName string
	PassWord string
	Email    string
	Role     RoleType
}

func CreateDataBase() {
	godotenv.Load()
	dsn := os.Getenv("DATABASE")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

}
