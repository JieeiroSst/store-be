package model

import "time"

type User struct {
	Id          string
	RoleId      string
	Username    string
	Passwrod    string
	Email       string
	Name        string
	Description string
	Address     string
	CreatedAt   time.Time
	Roles       []Role  `gorm:"foreignKey:Id;references:RoleId"`
	Sales       Sales   `gorm:"foreignKey:CustomerId"`
	Payment     Payment `gorm:"foreignKey:CustomerId"`
}

type Role struct {
	Id          string
	Title       string
	Description string
}

type Category struct {
	Id          string
	Name        string
	Description string
	Products    []Product `gorm:"foreignKey:CategoryId"`
}

type Product struct {
	Id          string
	Number      string
	Price       float32
	Type        string
	Description string
	CategoryId  string
	Discount    Discount `gorm:"foreignKey:ProductId"`
}

type Discount struct {
	Id          string
	Name        string
	Description string
	Type        string
	ProductId   string
}

type Sales struct {
	Id          string
	Amount      string
	Description string
	Type        string
	ProductId   string
	CustomerId  string
}

type Cart struct {
	Id       string
	Products []Product `gorm:"many2many:cart_product;"`
}

type Payment struct {
	Id          string
	CustomerId  string
	Amount      string
	Date        time.Time
	Description string
}
