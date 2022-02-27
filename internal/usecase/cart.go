package usecase

import (
	"github.com/JIeeiroSst/store/internal/repository"
	"github.com/JIeeiroSst/store/pkg/snowflake"
)

type Carts interface {
}

type CartUsecase struct {
	cartRepo  repository.Carts
	snowflake snowflake.SnowflakeData
}

func NewCartUsecase(cartRepo repository.Carts, snowflake snowflake.SnowflakeData) *CartUsecase {
	return &CartUsecase{
		cartRepo:  cartRepo,
		snowflake: snowflake,
	}
}
