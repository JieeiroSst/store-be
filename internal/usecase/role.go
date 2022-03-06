package usecase

import (
	"github.com/JIeeiroSst/store/internal/domain"
	"github.com/JIeeiroSst/store/internal/repository"
	"github.com/JIeeiroSst/store/model"
	"github.com/JIeeiroSst/store/pkg/snowflake"
)

type Roles interface {
	Create(model model.InputRole) error
	Update(id string, model model.InputRole) error
	Delete(id string) error
	RoleById(id string) (*domain.Role, error)
	RoleAll(pagination domain.Pagination) ([]domain.Role, error)
}

type RoleUsecase struct {
	roleRepo  repository.Roles
	snowflake snowflake.SnowflakeData
}

func NewRoleUsecase(roleRepo repository.Roles, snowflake snowflake.SnowflakeData) *RoleUsecase {
	return &RoleUsecase{
		roleRepo:  roleRepo,
		snowflake: snowflake,
	}
}

func (u *RoleUsecase) Create(model model.InputRole) error {
	role := domain.Role{
		Id:          u.snowflake.GearedID(),
		Title:       model.Title,
		Description: model.Description,
	}

	if err := u.roleRepo.Create(role); err != nil {
		return err
	}
	return nil
}

func (u *RoleUsecase) Update(id string, model model.InputRole) error {
	role := domain.Role{
		Title:       model.Title,
		Description: model.Description,
	}

	if err := u.roleRepo.Update(id, role); err != nil {
		return err
	}
	return nil
}

func (u *RoleUsecase) Delete(id string) error {
	if err := u.roleRepo.Delete(id); err != nil {
		return err
	}
	return nil
}

func (u *RoleUsecase) RoleById(id string) (*domain.Role, error) {
	role, err := u.roleRepo.RoleById(id)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (u *RoleUsecase) RoleAll(pagination domain.Pagination) ([]domain.Role, error) {
	roles, err := u.roleRepo.RoleAll(pagination)
	if err != nil {
		return nil, err
	}
	return roles, nil 
}
