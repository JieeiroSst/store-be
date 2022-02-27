package usecase

import (
	"github.com/JIeeiroSst/store/internal/repository"
	"github.com/JIeeiroSst/store/pkg/snowflake"
)

type Products interface {
}

type ProductUsecase struct {
	productRepo repository.Products
	snowflake   snowflake.SnowflakeData
}

func NewProductUsecase(productRepo repository.Products, snowflake snowflake.SnowflakeData) *ProductUsecase {
	return &ProductUsecase{
		productRepo: productRepo,
		snowflake:   snowflake,
	}
}
