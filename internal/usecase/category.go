package usecase

import (
	"github.com/JIeeiroSst/store/internal/repository"
	"github.com/JIeeiroSst/store/pkg/snowflake"
)

type Categories interface {
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
