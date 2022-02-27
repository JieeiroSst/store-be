package usecase

import (
	"github.com/JIeeiroSst/store/internal/repository"
	"github.com/JIeeiroSst/store/pkg/hash"
	"github.com/JIeeiroSst/store/pkg/jwt"
	"github.com/JIeeiroSst/store/pkg/snowflake"
)

type Usecase struct {
	Users      Users
	Roles      Roles
	Discounts  Discounts
	Payments   Payments
	Products   Products
	Categories Categories
	Sales      Sales
	Carts      Carts
}

type Dependency struct {
	Repos     *repository.Repositories
	snowflake snowflake.SnowflakeData
	hash      hash.Hash
	jwt       jwt.TokenUser
}

func NewUsecase(deps Dependency) *Usecase {
	userUsecase := NewUserUsecase(deps.Repos.Users, deps.snowflake, deps.hash, deps.jwt)
	roleUsecase := NewRoleUsecase(deps.Repos.Roles, deps.snowflake)
	discountUsecase := NewDiscountUsecase(deps.Repos.Discounts, deps.snowflake)
	paymentUsecase := NewPaymentUsecase(deps.Repos.Payments, deps.snowflake)
	productUsecase := NewProductUsecase(deps.Repos.Products, deps.snowflake)
	categoryUsecase := NewCategoryUsecase(deps.Repos.Categories, deps.snowflake)
	saleUsecase := NewSaleUsecase(deps.Repos.Sales, deps.snowflake)
	cartUsecase := NewCartUsecase(deps.Repos.Carts, deps.snowflake)

	return &Usecase{
		Users:      userUsecase,
		Roles:      roleUsecase,
		Discounts:  discountUsecase,
		Payments:   paymentUsecase,
		Products:   productUsecase,
		Categories: categoryUsecase,
		Sales:      saleUsecase,
		Carts:      cartUsecase,
	}
}
