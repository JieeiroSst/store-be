package model

type Cart struct {
	Id       string
	Products []Product `gorm:"many2many:cart_product;"`
}
