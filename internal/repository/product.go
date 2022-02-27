package repository

import (
	"errors"

	"github.com/JIeeiroSst/store/internal/domain"
	"gorm.io/gorm"
)

type Products interface {
	Create(product domain.Product) error
	Update(id string, product domain.Product) error
	Delete(id string) error
	ProductById(id string) (*domain.Product, error)
	Products(pagination domain.Pagination) ([]domain.Product, error)
}

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (r *ProductRepo) Create(product domain.Product) error {
	if err := r.db.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func (r *ProductRepo) Update(id string, product domain.Product) error {
	if err := r.db.Model(domain.Product{}).Where("id = ?", id).Updates(&product).Error; err != nil {
		return err
	}
	return nil
}

func (r *ProductRepo) Delete(id string) error {
	query := r.db.Delete(domain.Product{}, "id = ?", id)
	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected == 0 {
		return errors.New("delete product failed")
	}
	return nil
}

func (r *ProductRepo) ProductById(id string) (*domain.Product, error) {
	var product domain.Product
	query := r.db.Where("id = ?", id).Find(&product)
	if query.Error != nil {
		return nil, query.Error
	}
	if query.RowsAffected == 0 {
		return nil, errors.New("Not Found")
	}
	return &product, nil
}

func (r *ProductRepo) Products(pagination domain.Pagination) ([]domain.Product, error) {
	var products []domain.Product
	query := r.db.Limit(pagination.Limit).Offset(pagination.Offset).Find(&products)
	if query.Error != nil {
		return nil, query.Error
	}
	if query.RowsAffected == 0 {
		return nil, errors.New("Not Found")
	}
	return products, nil
}
