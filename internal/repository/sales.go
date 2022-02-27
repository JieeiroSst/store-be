package repository

import (
	"errors"

	"github.com/JIeeiroSst/store/internal/domain"
	"gorm.io/gorm"
)

type Sales interface {
	Create(sale domain.Sale) error
	Update(id string, sale domain.Sale) error
	Delete(id string) error
	Sales(pagination domain.Pagination) ([]domain.Sale, error)
	SaleById(id string) (*domain.Sale, error)
}

type SaleRepo struct {
	db *gorm.DB
}

func NewSaleRepo(db *gorm.DB) *SaleRepo {
	return &SaleRepo{
		db: db,
	}
}

func (r *SaleRepo) Create(sale domain.Sale) error {
	if err := r.db.Create(&sale).Error; err != nil {
		return err
	}
	return nil
}

func (r *SaleRepo) Update(id string, sale domain.Sale) error {
	if err := r.db.Model(domain.Sale{}).Where("id = ?", id).Updates(&sale).Error; err != nil {
		return err
	}
	return nil
}

func (r *SaleRepo) Delete(id string) error {
	query := r.db.Delete(domain.Sale{}, "id = ?", id)
	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected == 0 {
		return errors.New("delete sale failed")
	}
	return nil
}

func (r *SaleRepo) Sales(pagination domain.Pagination) ([]domain.Sale, error) {
	var sales []domain.Sale
	query := r.db.Limit(pagination.Limit).Offset(pagination.Offset).Find(&sales)
	if query.Error != nil {
		return nil, query.Error
	}
	if query.RowsAffected == 0 {
		return nil, errors.New("Not Found")
	}
	return sales, nil
}

func (r *SaleRepo) SaleById(id string) (*domain.Sale, error) {
	var sale domain.Sale
	query := r.db.Where("id = ?", id).Find(&sale)
	if query.Error != nil {
		return nil, query.Error
	}
	if query.RowsAffected == 0 {
		return nil, errors.New("Not Found")
	}
	return &sale, nil
}
