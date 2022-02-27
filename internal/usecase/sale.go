package usecase

import (
	"github.com/JIeeiroSst/store/internal/repository"
	"github.com/JIeeiroSst/store/pkg/snowflake"
)

type Sales interface{}

type SaleUsecase struct {
	saleRepo  repository.Sales
	snowflake snowflake.SnowflakeData
}

func NewSaleUsecase(saleRepo repository.Sales, snowflake snowflake.SnowflakeData) *SaleUsecase {
	return &SaleUsecase{
		saleRepo:  saleRepo,
		snowflake: snowflake,
	}
}
