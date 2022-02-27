package usecase

import (
	"github.com/JIeeiroSst/store/internal/repository"
	"github.com/JIeeiroSst/store/pkg/snowflake"
)

type Discounts interface{}

type DiscountUsecase struct {
	discountRepo repository.Discounts
	snowflake    snowflake.SnowflakeData
}

func NewDiscountUsecase(discountRepo repository.Discounts, snowflake snowflake.SnowflakeData) *DiscountUsecase {
	return &DiscountUsecase{
		discountRepo: discountRepo,
		snowflake:    snowflake,
	}
}
