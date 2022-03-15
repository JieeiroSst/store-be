package model

type Category struct {
	Id          string
	Name        string
	Description string
	Products    []Product `gorm:"foreignKey:CategoryId"`
}
