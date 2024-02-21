package databases

import (
	"time"

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
	Order    []Order   `gorm:"foreignKey:UserId;refernces:Id"`
}

type Address struct {
	gorm.Model
	ID        int `gorm:"AUTO_INCREMENT"`
	FirstName string
	LastName  string
	Address   string
	Phone     string
	UserId    int
	Order     []Order `gorm:"foreignKey:AddressId;refernces:Id"`
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
	OrderItem    []OrderItem    `gorm:"foreignKey:ProductId;refernces:Id"`
}

type ImageProduct struct {
	gorm.Model
	ID        int `gorm:"AUTO_INCREMENT"`
	ImageUrl  string
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

type PaymentStatus string

const (
	PENDING = "PENDING"
	SUCCESS = "SUCCESS"
)

type Order struct {
	gorm.Model
	ID            int `gorm:"AUTO_INCREMENT"`
	TotalPrice    float32
	QrImageUrl    string
	SlipUrl       string
	PaymentDate   time.Time
	PaymentStatus PaymentStatus
	UserId        int
	AddressId     int
	OrderItem     []OrderItem `gorm:"foreignKey:OrderId;refernces:Id"`
}

type OrderItem struct {
	gorm.Model
	ID         int `gorm:"AUTO_INCREMENT"`
	Quantity   int
	TotalPrice int
	ProductId  int
	OrderId    int
}
