package usecase

import (
	"github.com/JIeeiroSst/store/internal/domain"
	"github.com/JIeeiroSst/store/internal/repository"
	"github.com/JIeeiroSst/store/pkg/snowflake"
)

type Categories interface {
	Create(category domain.Category) error
	Update(id string, category domain.Category) error
	Delete(id string) error
	Categories(pagination domain.Pagination) ([]domain.Category, error)
	CategoryById(id string) (*domain.Category, error)
}

type CategoryUsecase struct {
	categoryRepo repository.Categories
	snowflake    snowflake.SnowflakeData
}

func NewCategoryUsecase(categoryRepo repository.Categories, snowflake snowflake.SnowflakeData) *CategoryUsecase {
	return &CategoryUsecase{
		categoryRepo: categoryRepo,
		snowflake:    snowflake,
	}
}

func (u *CategoryUsecase) Create(category domain.Category) error {
	if err := u.categoryRepo.Create(category); err != nil {
		return err
	}
	return nil
}

func (u *CategoryUsecase) Update(id string, category domain.Category) error {
	if err := u.categoryRepo.Update(id, category); err != nil {
		return err
	}
	return nil
}

func (u *CategoryUsecase) Delete(id string) error {
	if err := u.categoryRepo.Delete(id); err != nil {
		return err
	}
	return nil
}

func (u *CategoryUsecase) Categories(pagination domain.Pagination) ([]domain.Category, error) {
	categories, err := u.categoryRepo.Categories(pagination)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (u *CategoryUsecase) CategoryById(id string) (*domain.Category, error) {
	category, err := u.categoryRepo.CategoryById(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}
