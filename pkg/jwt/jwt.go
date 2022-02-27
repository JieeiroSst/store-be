package jwt

import (
	"time"

	"github.com/JIeeiroSst/store/model"
	"github.com/dgrijalva/jwt-go"
)

type tokenUser struct {
	serect string
}

type TokenUser interface {
	GenerateToken(token model.Token) (*model.ResultToken, error)
	ParseToken(tokenStr string) (*model.Token, error)
}

func NewTokenUser(serect string) TokenUser {
	return &tokenUser{
		serect: serect,
	}
}

func (t *tokenUser) GenerateToken(token model.Token) (*model.ResultToken, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["username"] = token.Username
	atClaims["roleId"] = token.RoleId
	atClaims["id"] = token.Id
	atClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	tokenStr, err := at.SignedString([]byte(t.serect))
	if err != nil {
		return nil, err
	}
	return &model.ResultToken{
		Token: tokenStr,
	}, nil
}

func (t *tokenUser) ParseToken(tokenStr string) (*model.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.serect), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
	} else {
		return nil, err
	}

	return &model.Token{
		Id:       claims["id"].(string),
		Username: claims["username"].(string),
		RoleId:   claims["roleId"].(string),
	}, nil
}
