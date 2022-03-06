package jwt

import (
	"time"

	"github.com/JIeeiroSst/store/model"
	"github.com/JIeeiroSst/store/pkg/snowflake"
	"github.com/dgrijalva/jwt-go"
)

type tokenUser struct {
	accessSerect  string
	refreshSerect string
	snowflake     snowflake.SnowflakeData
}

type TokenUser interface {
	CreateToken(token model.Token) (*model.TokenDetails, error)
	ParseToken(tokenStr string) (*model.Token, error)
	ParseRefreshToken(tokenStr string) (*model.TokenDetails, error)
	RegenerateAccessToken(token model.Token, accessUuid string) (string, error)
}

func NewTokenUser(accessSerect string, refreshSerect string, snowflake snowflake.SnowflakeData) TokenUser {
	return &tokenUser{
		accessSerect:  accessSerect,
		refreshSerect: refreshSerect,
		snowflake:     snowflake,
	}
}

func (t *tokenUser) ParseToken(tokenStr string) (*model.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.accessSerect), nil
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

func (t *tokenUser) ParseRefreshToken(tokenStr string) (*model.TokenDetails, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.refreshSerect), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
	} else {
		return nil, err
	}

	return &model.TokenDetails{
		RefreshUuid: claims["refresh_uuid"].(string),
	}, nil
}

func (t *tokenUser) CreateToken(token model.Token) (*model.TokenDetails, error) {
	td := &model.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()

	td.AccessUuid = t.snowflake.GearedID()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = t.snowflake.GearedID()

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["exp"] = td.AtExpires
	atClaims["authorized"] = true
	atClaims["username"] = token.Username
	atClaims["roleId"] = token.RoleId
	atClaims["id"] = token.Id
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(t.accessSerect))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["exp"] = td.RtExpires
	rtClaims["authorized"] = true
	rtClaims["username"] = token.Username
	rtClaims["roleId"] = token.RoleId
	rtClaims["id"] = token.Id
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(t.refreshSerect))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func (t *tokenUser) RegenerateAccessToken(token model.Token, accessUuid string) (string, error) {
	atExpires := time.Now().Add(time.Minute * 15).Unix()
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = accessUuid
	atClaims["exp"] = atExpires
	atClaims["authorized"] = true
	atClaims["username"] = token.Username
	atClaims["roleId"] = token.RoleId
	atClaims["id"] = token.Id
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	accessToken, err := at.SignedString([]byte(t.accessSerect))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}
