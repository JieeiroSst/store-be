package repository

import (
	"errors"

	"github.com/JIeeiroSst/store/internal/domain"
	"gorm.io/gorm"
)

type Categories interface {
	Create(category domain.Category) error
	Update(id string, category domain.Category) error
	Delete(id string) error
	Categories(pagination domain.Pagination) ([]domain.Category, error)
	CategoryById(id string) (*domain.Category, error)
}

type CategoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{
		db: db,
	}
}

func (r *CategoryRepo) Create(category domain.Category) error {
	if err := r.db.Create(&category).Error; err != nil {
		return err
	}
	return nil
}

func (r *CategoryRepo) Update(id string, category domain.Category) error {
	if err := r.db.Model(domain.Category{}).Where("id = ?", id).Updates(&category).Error; err != nil {
		return err
	}
	return nil
}

func (r *CategoryRepo) Delete(id string) error {
	query := r.db.Delete(domain.Category{}, "id = ?", id)
	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected == 0 {
		return errors.New("delete category failed")
	}
	return nil
}

func (r *CategoryRepo) Categories(pagination domain.Pagination) ([]domain.Category, error) {
	var categories []domain.Category
	query := r.db.Limit(pagination.Limit).Offset(pagination.Offset).Find(&categories)
	if query.Error != nil {
		return nil, query.Error
	}
	if query.RowsAffected == 0 {
		return nil, errors.New("Not Found")
	}
	return categories, nil
}

func (r *CategoryRepo) CategoryById(id string) (*domain.Category, error) {
	var category domain.Category
	query := r.db.Where("id = ?", id).Find(&category)
	if query.Error != nil {
		return nil, query.Error
	}
	if query.RowsAffected == 0 {
		return nil, errors.New("Not Found")
	}
	return &category, nil
}
