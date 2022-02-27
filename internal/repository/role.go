package repository

import (
	"errors"

	"github.com/JIeeiroSst/store/internal/domain"
	"gorm.io/gorm"
)

type Roles interface {
	Create(role domain.Role) error
	Update(id string, role domain.Role) error
	Delete(id string) error
	RoleById(id string) (*domain.Role, error)
	RoleAll(pagination domain.Pagination) ([]domain.Role, error)
}

type RoleRepo struct {
	db *gorm.DB
}

func NewRoleRepo(db *gorm.DB) *RoleRepo {
	return &RoleRepo{
		db: db,
	}
}

func (r *RoleRepo) Create(role domain.Role) error {
	if err := r.db.Create(&role).Error; err != nil {
		return err
	}
	return nil
}

func (r *RoleRepo) Update(id string, role domain.Role) error {
	if err := r.db.Model(domain.Role{}).Where("id = ?", id).Updates(&role).Error; err != nil {
		return err
	}
	return nil
}

func (r *RoleRepo) Delete(id string) error {
	query := r.db.Delete(domain.Role{}, "id = ?", id)
	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected == 0 {
		return errors.New("delete task failed")
	}
	return nil
}

func (r *RoleRepo) RoleById(id string) (*domain.Role, error) {
	var role domain.Role
	query := r.db.Where("id = ?", id).Find(&role)
	if query.Error != nil {
		return nil, query.Error
	}

	if query.RowsAffected == 0 {
		return nil, errors.New("")
	}

	return &role, nil
}

func (r *RoleRepo) RoleAll(pagination domain.Pagination) ([]domain.Role, error) {
	var roles []domain.Role
	query := r.db.Limit(pagination.Limit).Offset(pagination.Offset).Find(&roles)
	if query.Error != nil {
		return nil, query.Error
	}

	if query.RowsAffected == 0 {
		return nil, errors.New("")
	}

	return roles, nil
}
