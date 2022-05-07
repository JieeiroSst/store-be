package model

type Cart struct {
	Id       string
	UserId   string
	Products []Product `gorm:"many2many:cart_product;"`
}
