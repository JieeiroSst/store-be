package usecase

import (
	"github.com/JIeeiroSst/store/internal/repository"
	"github.com/JIeeiroSst/store/pkg/snowflake"
)

type Roles interface {
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
