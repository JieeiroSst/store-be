package repository

import (
	"context"

	"github.com/JIeeiroSst/store/internal/domain"
	"gorm.io/gorm"
)

type Medias interface {
	SaveMedia(ctx context.Context, args domain.Media) error
	UpdateMedia(ctx context.Context, id int, args domain.Media) error
}

type MediaRepo struct {
	db *gorm.DB
}

func NewMediaRepo(db *gorm.DB) *MediaRepo {
	return &MediaRepo{
		db: db,
	}
}

func (r *MediaRepo) SaveMedia(ctx context.Context, args domain.Media) error {
	if err := r.db.Create(&args).Error; err != nil {
		return err
	}
	return nil
}

func (r *MediaRepo) UpdateMedia(ctx context.Context, id int, args domain.Media) error {
	if err := r.db.Model(domain.Media{}).Where("id = ?", id).Updates(&args).Error; err != nil {
		return err
	}
	return nil
}
