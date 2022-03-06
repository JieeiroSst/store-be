package repository

import "gorm.io/gorm"

type Repositories struct {
	Users       Users
	Roles       Roles
	Discounts   Discounts
	Payments    Payments
	Products    Products
	Categories  Categories
	Sales       Sales
	Carts       Carts
	Permissions Permissions
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Users:       NewUserRepo(db),
		Roles:       NewRoleRepo(db),
		Discounts:   NewDiscountRepo(db),
		Payments:    NewPaymentRepo(db),
		Products:    NewProductRepo(db),
		Categories:  NewCategoryRepo(db),
		Sales:       NewSaleRepo(db),
		Carts:       NewCartRepo(db),
		Permissions: NewPermissionRepo(db),
	}
}
