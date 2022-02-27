package usecase

import (
	"github.com/JIeeiroSst/store/internal/repository"
	"github.com/JIeeiroSst/store/pkg/snowflake"
)

type Payments interface{}

type PaymentUsecase struct {
	paymentRepo repository.Payments
	snowflake   snowflake.SnowflakeData
}

func NewPaymentUsecase(paymentRepo repository.Payments, snowflake snowflake.SnowflakeData) *PaymentUsecase {
	return &PaymentUsecase{
		paymentRepo: paymentRepo,
		snowflake:   snowflake,
	}
}
