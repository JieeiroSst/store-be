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
	Tokens     Tokens
}

type Dependency struct {
	Repos     *repository.Repositories
	Snowflake snowflake.SnowflakeData
	Hash      hash.Hash
	Jwt       jwt.TokenUser
}

func NewUsecase(deps Dependency) *Usecase {
	userUsecase := NewUserUsecase(deps.Repos.Users, deps.Snowflake, deps.Hash, deps.Jwt)
	roleUsecase := NewRoleUsecase(deps.Repos.Roles, deps.Snowflake)
	discountUsecase := NewDiscountUsecase(deps.Repos.Discounts, deps.Snowflake)
	paymentUsecase := NewPaymentUsecase(deps.Repos.Payments, deps.Snowflake)
	productUsecase := NewProductUsecase(deps.Repos.Products, deps.Snowflake)
	categoryUsecase := NewCategoryUsecase(deps.Repos.Categories, deps.Snowflake)
	saleUsecase := NewSaleUsecase(deps.Repos.Sales, deps.Snowflake)
	cartUsecase := NewCartUsecase(deps.Repos.Carts, deps.Snowflake)
	tokenUsecase := NewTokenUsecase(deps.Jwt)

	return &Usecase{
		Users:      userUsecase,
		Roles:      roleUsecase,
		Discounts:  discountUsecase,
		Payments:   paymentUsecase,
		Products:   productUsecase,
		Categories: categoryUsecase,
		Sales:      saleUsecase,
		Carts:      cartUsecase,
		Tokens:     tokenUsecase,
	}
}
