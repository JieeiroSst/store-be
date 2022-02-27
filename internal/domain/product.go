package domain

type Product struct {
	Id          string
	Number      string
	Price       float32
	Type        string
	Description string
	CategoryId  string
	Discount    Discount `gorm:"foreignKey:ProductId"`
	Carts       []Cart   `gorm:"many2many:cart_product;"`
}
