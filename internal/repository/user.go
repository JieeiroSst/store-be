package repository

import (
	"errors"
	"fmt"

	"github.com/JIeeiroSst/store/internal/domain"
	"gorm.io/gorm"
)

type Users interface {
	CheckAccount(user domain.User) (*domain.Token, error)
	CheckAccountExists(user domain.User) error
	Create(user domain.User) error
	Update(id string, user domain.User) error
	UserById(id string) (*domain.User, error)
	UserByAll(pagination domain.Pagination) ([]domain.User, error)
	LockUser(id string) error
	UnLockLockUser(id string) error
	IsUserLock(id string) error
	
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Create(user domain.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) Update(id string, user domain.User) error {
	if err := r.db.Model(domain.User{}).Where("id = ?", id).Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) UserById(id string) (*domain.User, error) {
	var user domain.User
	query := r.db.Where("id = ?", id).Find(&user)
	if query.Error != nil {
		return nil, query.Error
	}
	if query.RowsAffected == 0 {
		return nil, errors.New("")
	}
	return &user, nil
}

func (r *UserRepo) UserByAll(pagination domain.Pagination) ([]domain.User, error) {
	var users []domain.User
	query := r.db.Limit(pagination.Limit).Offset(pagination.Offset).Find(&users)
	if query.Error != nil {
		return nil, query.Error
	}
	if query.RowsAffected == 0 {
		return nil, errors.New("")
	}
	return users, nil
}

func (r *UserRepo) LockUser(id string) error {
	if err := r.db.Model(domain.User{}).Where("id = ?", id).
		Updates(domain.User{Lock: -1}).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) UnLockLockUser(id string) error {
	if err := r.db.Model(domain.User{}).Where("id = ?", id).
		Updates(domain.User{Lock: 1}).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) CheckAccount(user domain.User) (*domain.Token, error) {
	var result domain.User
	query := r.db.Where("username = ?", user.Username).Limit(1).Find(&result)

	if query.Error != nil {
		return nil, errors.New("query error")
	}

	if query.RowsAffected == 0 {
		return nil, errors.New("user does not exist")
	}
	return &domain.Token{
		Id:       result.Id,
		RoleId:   result.RoleId,
		Username: result.Username,
		Password: result.Password,
	}, nil
}
func (r *UserRepo) CheckAccountExists(user domain.User) error {
	var result domain.User
	query := r.db.Where("username = ?", user.Username).Limit(1).Find(&result)
	if query.Error != nil {
		return errors.New("query error")
	}

	if query.RowsAffected == 1 {
		return errors.New("user does exist")
	}
	return nil
}

func (r *UserRepo) IsUserLock(id string) error {
	var user domain.User
	query := r.db.Where("id = ?", id).Find(&user)
	if query.Error != nil {
		return query.Error
	}

	if query.RowsAffected == 0 {
		return fmt.Errorf("Not Found User by id %s", id)
	}

	if user.Lock == -1 {
		return errors.New("User is lock")
	}

	return nil
}
