package service

import (
	"dtyTrade/config"
	"dtyTrade/models"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func RefreshAccessToken(id int64, email string) (string, error) {

	claim := jwt.MapClaims{
		"id":        id,
		"email":     email,
		"expiredAt": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(config.Conf.JwtSecret))
}

func CheckToken(tokenStr string) (*models.User, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Conf.JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("cannot convert claim to MapClaims")
	}
	if !token.Valid {
		return nil, errors.New("token is invalid")
	}

	email := claim["email"].(string)

	user, err := GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("bad token")
	}
	return user, nil
}
