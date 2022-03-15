package usecase

import (
	"github.com/JIeeiroSst/store/model"
	"github.com/JIeeiroSst/store/pkg/jwt"
)

type Tokens interface {
	RegenerateAccessToken(token model.Token, accessUuid string) (string, error)
}
type TokenUsecase struct {
	jwt jwt.TokenUser
}

func NewTokenUsecase(jwt jwt.TokenUser) *TokenUsecase {
	return &TokenUsecase{
		jwt: jwt,
	}
}

func (u *TokenUsecase) RegenerateAccessToken(token model.Token, accessUuid string) (string, error) {
	accessToken, err := u.jwt.RegenerateAccessToken(token, accessUuid)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}
