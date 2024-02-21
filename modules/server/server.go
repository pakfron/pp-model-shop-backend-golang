package server

import (
	"fmt"
	"log"
	"os"
	databases "pp-model-shop-backend/pkg/database"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Instance *gorm.DB
)

func MigrateDatabase() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	godotenv.Load("../.env")
	dsn := os.Getenv("DATABASE")
	fmt.Printf("%v", &dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&databases.User{}, &databases.Address{}, &databases.Address{}, &databases.ImageProduct{},
		&databases.Cart{}, &databases.Order{}, &databases.OrderItem{})
	defer CloseDB(db)

}

func GetDB() *gorm.DB {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	godotenv.Load("../.env")
	dsn := os.Getenv("DATABASE")
	fmt.Printf("%v", &dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func CloseDB(db *gorm.DB) error {

	InstantDB, _ := db.DB()
	return InstantDB.Close()
}
