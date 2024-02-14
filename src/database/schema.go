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
	ID       int `gorm:"AUTO_INCREMENT"`
	UserName string
	PassWord string
	Email    string
	Role     RoleType  `gorm:"default:USER"`
	Address  []Address `gorm:"foreignKey:UserId;refernces:Id"`
	Cart     []Cart    `gorm:"foreignKey:UserId;refernces:Id"`
}

type Address struct {
	gorm.Model
	ID        int `gorm:"AUTO_INCREMENT"`
	FirstName string
	LastName  string
	Address   string
	Phone     string
	UserId    int
}

type Type string

const (
	Nendoroid = "Nendoroid"
	Figma     = "Figma"
	Figure    = "Figure"
)

type Product struct {
	gorm.Model
	ID           int `gorm:"AUTO_INCREMENT"`
	Name         string
	Series       string
	Detail       string
	Price        float32
	Type         Type
	IsActive     bool           `gorm:"default:true"`
	ImageProduct []ImageProduct `gorm:"foreignKey:ProductId;refernces:Id"`
	Cart         []Cart         `gorm:"foreignKey:ProductId;refernces:Id"`
}

type ImageProduct struct {
	gorm.Model
	ID        int `gorm:"AUTO_INCREMENT"`
	imageUrl  string
	ProductId int
}

type Cart struct {
	gorm.Model
	ID         int `gorm:"AUTO_INCREMENT"`
	TotalPrice float32
	Quantity   int
	UserId     int
	ProductId  int
}

var (
	Instance *gorm.DB
)

func CreateDataBase() {
	godotenv.Load()
	dsn := os.Getenv("DATABASE")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{}, &Address{}, &Product{}, &ImageProduct{}, &Cart{})
	Instance = db

}
