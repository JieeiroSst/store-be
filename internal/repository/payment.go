package repository

import (
	"errors"

	"github.com/JIeeiroSst/store/internal/domain"
	"gorm.io/gorm"
)

type Payments interface {
	Create(payment domain.Payment) error
	Update(id string, payment domain.Payment) error
	Delete(id string) error
	Payments(pagination domain.Pagination) ([]domain.Payment, error)
	PaymentById(id string) (*domain.Payment, error)
}

type PaymentRepo struct {
	db *gorm.DB
}

func NewPaymentRepo(db *gorm.DB) *PaymentRepo {
	return &PaymentRepo{
		db: db,
	}
}

func (r *PaymentRepo) Create(payment domain.Payment) error {
	if err := r.db.Create(&payment).Error; err != nil {
		return err
	}
	return nil
}

func (r *PaymentRepo) Update(id string, payment domain.Payment) error {
	if err := r.db.Model(domain.Payment{}).Where("id = ?", id).Updates(&payment).Error; err != nil {
		return err
	}
	return nil
}

func (r *PaymentRepo) Delete(id string) error {
	query := r.db.Delete(domain.Payment{}, "id = ?", id)
	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected == 0 {
		return errors.New("delete payment failed")
	}
	return nil
}

func (r *PaymentRepo) Payments(pagination domain.Pagination) ([]domain.Payment, error) {
	var payments []domain.Payment
	query := r.db.Limit(pagination.Limit).Offset(pagination.Offset).Find(&payments)
	if query.Error != nil {
		return nil, query.Error
	}
	if query.RowsAffected == 0 {
		return nil, errors.New("Not Found")
	}
	return payments, nil
}

func (r *PaymentRepo) PaymentById(id string) (*domain.Payment, error) {
	var payment domain.Payment
	query := r.db.Where("id = ?", id).Find(&payment)
	if query.Error != nil {
		return nil, query.Error
	}
	if query.RowsAffected == 0 {
		return nil, errors.New("Not Found")
	}
	return &payment, nil
}
