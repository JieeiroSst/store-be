package repository

import (
	"errors"

	"github.com/JIeeiroSst/store/internal/domain"
	"gorm.io/gorm"
)

type Permissions interface {
	Create(permission domain.Permission) error
	Update(id string, permission domain.Permission) error
	Delete(id string) error
	PermissionById(id string) (*domain.Permission, error)
	Permissions(pagination domain.Pagination) ([]domain.Permission, error)
}

type PermissionRepo struct {
	db *gorm.DB
}

func NewPermissionRepo(db *gorm.DB) *PermissionRepo {
	return &PermissionRepo{
		db: db,
	}
}

func (r *PermissionRepo) Create(permission domain.Permission) error {
	if err := r.db.Create(&permission).Error; err != nil {
		return err
	}
	return nil
}

func (r *PermissionRepo) Update(id string, permission domain.Permission) error {
	if err := r.db.Model(domain.Permission{}).Where("id = ?", id).Updates(&permission).Error; err != nil {
		return err
	}
	return nil
}

func (r *PermissionRepo) Delete(id string) error {
	query := r.db.Delete(domain.Permission{}, "id = ?", id)
	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected == 0 {
		return errors.New("delete task failed")
	}
	return nil
}

func (r *PermissionRepo) PermissionById(id string) (*domain.Permission, error) {
	var permission domain.Permission
	query := r.db.Where("id = ?", id).Find(&permission)
	if query.Error != nil {
		return nil, query.Error
	}
	if query.RowsAffected == 0 {
		return nil, errors.New("Not Found")
	}
	return &permission, nil
}

func (r *PermissionRepo) Permissions(pagination domain.Pagination) ([]domain.Permission, error) {
	var permissions []domain.Permission
	query := r.db.Limit(pagination.Limit).Offset(pagination.Offset).Find(&permissions)
	if query.Error != nil {
		return nil, query.Error
	}
	if query.RowsAffected == 0 {
		return nil, errors.New("Not Found")
	}
	return permissions, nil
}
