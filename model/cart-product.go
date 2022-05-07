package model

type CartProduct struct {
	ProductId string `gorm:"primaryKey" column:"product_id"`
	CartId    string `gorm:"primaryKey" column:"cart_id"`
	Active    bool
}
