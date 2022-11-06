package service

import (
	"errors"
	"time"

	"github.com/cleysonph/hyperprof/config"
	"github.com/cleysonph/hyperprof/internal/database"
	"github.com/golang-jwt/jwt/v4"
)

func generateAccessToken(sub string) (string, error) {
	return generateToken(sub, config.TokenSecret, time.Duration(config.TokenDuration)*time.Second)
}

func generateRefreshToken(sub string) (string, error) {
	return generateToken(sub, config.RefreshSecret, time.Duration(config.RefreshDuration)*time.Second)
}

func generateToken(sub, secret string, exp time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = sub
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(exp).Unix()
	return token.SignedString([]byte(secret))
}

func getSubFromAccessToken(token string) (string, error) {
	return getSubFromToken(token, config.TokenSecret)
}

func getSubFromRefreshToken(token string) (string, error) {
	return getSubFromToken(token, config.RefreshSecret)
}

func getSubFromToken(token, secret string) (string, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}

	err = validateToken(token, secret)
	if err != nil {
		return "", err
	}

	return claims["sub"].(string), nil
}

func validateToken(token, secret string) error {
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return err
	}

	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			return err
		}
	}

	if database.ExistsInvalidatedTokenByToken(token) {
		return errors.New("token has been invalidated")
	}

	return nil
}

func invalidateTokens(tokens ...string) {
	for _, token := range tokens {
		database.CreateInvalidatedToken(token)
	}
}
