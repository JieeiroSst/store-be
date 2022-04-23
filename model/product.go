package model

type Product struct {
	Id          string
	Number      string
	Price       float32
	Type        string
	Description string
	CategoryId  string
	Discount    Discount `gorm:"foreignKey:ProductId"`
}

type InputProduct struct {
	Number      string
	Price       float32
	Type        string
	Description string
	CategoryId  string
}
