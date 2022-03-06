package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/JIeeiroSst/store/model"
	"github.com/JIeeiroSst/store/pkg/hash"
	"github.com/JIeeiroSst/store/pkg/jwt"
	"github.com/JIeeiroSst/store/pkg/redis"

	"github.com/JIeeiroSst/store/internal/domain"
	"github.com/JIeeiroSst/store/internal/repository"
	"github.com/JIeeiroSst/store/pkg/snowflake"
)

type Users interface {
	Login(user domain.User) (*model.TokenDetails, error)
	SignUp(user domain.User) error
	Update(id string, input domain.InputUser) error
	UserById(id string) (*domain.User, error)
	UserByAll(pagination domain.Pagination) ([]domain.User, error)
	LockUser(id string) error
}

type UserUsecase struct {
	jwt       jwt.TokenUser
	hash      hash.Hash
	userRepo  repository.Users
	snowflake snowflake.SnowflakeData
	cache     redis.RedisDB
}

func NewUserUsecase(userRepo repository.Users, snowflake snowflake.SnowflakeData,
	hash hash.Hash, jwt jwt.TokenUser) *UserUsecase {
	return &UserUsecase{
		jwt:       jwt,
		hash:      hash,
		userRepo:  userRepo,
		snowflake: snowflake,
	}
}

func (u *UserUsecase) Login(user domain.User) (*model.TokenDetails, error) {
	token, err := u.userRepo.CheckAccount(user)
	if err != nil {
		return nil, err
	}

	if err := u.userRepo.IsUserLock(token.Id); err != nil {
		return nil, err
	}

	if err := u.hash.CheckPassword(user.Password, token.Password); err != nil {
		return nil, err
	}

	tokens := model.Token{
		Id:       token.Id,
		Username: token.Username,
		RoleId:   token.RoleId,
	}

	tokenString, err := u.jwt.CreateToken(tokens)
	if err != nil {
		return nil, err
	}

	if err := u.cache.CreateAuth(context.Background(), token.Id, tokenString); err != nil {
		return nil, err
	}

	return tokenString, nil
}

func (u *UserUsecase) SignUp(user domain.User) error {
	check := u.userRepo.CheckAccountExists(user)
	if check != nil {
		return errors.New("user already exists")
	}
	hashPassword, err := u.hash.HashPassword(user.Password)
	if err != nil {
		return errors.New("password failed")
	}
	account := domain.User{
		Id:        u.snowflake.GearedID(),
		Username:  user.Username,
		Password:  hashPassword,
		CreatedAt: time.Now(),
		Lock:      -1,
	}
	if err = u.userRepo.Create(account); err != nil {
		return errors.New("create failed")
	}
	return nil
}

func (u *UserUsecase) Update(id string, input domain.InputUser) error {
	user := domain.User{
		RoleId:      input.RoleId,
		Email:       input.Email,
		Name:        input.Name,
		Description: input.Description,
		Address:     input.Address,
	}
	if err := u.userRepo.Update(id, user); err != nil {
		return err
	}

	return nil
}

func (u *UserUsecase) UserById(id string) (*domain.User, error) {
	user, err := u.userRepo.UserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (u *UserUsecase) UserByAll(pagination domain.Pagination) ([]domain.User, error) {
	users, err := u.userRepo.UserByAll(pagination)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserUsecase) LockUser(id string) error {
	if err := r.userRepo.LockUser(id); err != nil {
		return err
	}
	return nil
}
