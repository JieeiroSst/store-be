package usecase

import (
	"github.com/JIeeiroSst/store/internal/domain"
	"github.com/JIeeiroSst/store/internal/repository"
	"github.com/JIeeiroSst/store/pkg/snowflake"
)

type Payments interface {
	Create(payment domain.Payment) error
	Update(id string, payment domain.Payment) error
	Delete(id string) error
	Payments(pagination domain.Pagination) ([]domain.Payment, error)
	PaymentById(id string) (*domain.Payment, error)
}

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

func (u *PaymentUsecase) Create(payment domain.Payment) error {
	args := domain.Payment{
		Id:          u.snowflake.GearedID(),
		CustomerId:  payment.CustomerId,
		Amount:      payment.Amount,
		Date:        payment.Date,
		Description: payment.Description,
	}
	if err := u.paymentRepo.Create(args); err != nil {
		return err
	}
	return nil
}

func (u *PaymentUsecase) Update(id string, payment domain.Payment) error {
	if err := u.paymentRepo.Update(id, payment); err != nil {
		return err
	}
	return nil
}

func (u *PaymentUsecase) Delete(id string) error {
	if err := u.paymentRepo.Delete(id); err != nil {
		return err
	}
	return nil
}

func (u *PaymentUsecase) Payments(pagination domain.Pagination) ([]domain.Payment, error) {
	payments, err := u.paymentRepo.Payments(pagination)
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func (u *PaymentUsecase) PaymentById(id string) (*domain.Payment, error) {
	payment, err := u.paymentRepo.PaymentById(id)
	if err != nil {
		return nil, err
	}
	return payment, nil
}
