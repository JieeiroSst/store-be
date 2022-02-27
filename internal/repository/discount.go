package repository

import (
	"errors"

	"github.com/JIeeiroSst/store/internal/domain"
	"gorm.io/gorm"
)

type Discounts interface {
	Create(discount domain.Discount) error
	Update(id string, discount domain.Discount) error
	Delete(id string) error
	Discounts(pagination domain.Pagination) ([]domain.Discount, error)
	DiscountById(id string) (*domain.Discount, error)
}

type DiscountRepo struct {
	db *gorm.DB
}

func NewDiscountRepo(db *gorm.DB) *DiscountRepo {
	return &DiscountRepo{
		db: db,
	}
}

func (r *DiscountRepo) Create(discount domain.Discount) error {
	if err := r.db.Create(&discount).Error; err != nil {
		return err
	}
	return nil
}

func (r *DiscountRepo) Update(id string, discount domain.Discount) error {
	if err := r.db.Model(domain.Discount{}).Where("id = ?", id).Updates(&discount).Error; err != nil {
		return err
	}
	return nil
}

func (r *DiscountRepo) Delete(id string) error {
	query := r.db.Delete(domain.Discount{}, "id = ?", id)
	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected == 0 {
		return errors.New("delete discount failed")
	}
	return nil
}

func (r *DiscountRepo) Discounts(pagination domain.Pagination) ([]domain.Discount, error) {
	var discounts []domain.Discount
	query := r.db.Limit(pagination.Limit).Offset(pagination.Offset).Find(&discounts)
	if query.Error != nil {
		return nil, query.Error
	}
	if query.RowsAffected == 0 {
		return nil, errors.New("Not found discount")
	}
	return discounts, nil
}

func (r *DiscountRepo) DiscountById(id string) (*domain.Discount, error) {
	var discount domain.Discount
	query := r.db.Where("id = ?", id).Find(&discount)
	if query.Error != nil {
		return nil, query.Error
	}
	if query.RowsAffected == 0 {
		return nil, errors.New("Not found discount")
	}
	return &discount, nil
}
