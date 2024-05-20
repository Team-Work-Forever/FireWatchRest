package jwtService

import (
	"time"

	"github.com/Team-Work-Forever/FireWatchRest/config"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJwtToken(duration time.Duration) (string, error) {
	env := config.GetCofig()

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = "user_id"
	claims["email"] = "email@gmail.com"
	claims["role"] = "admin"
	claims["iss"] = env.FIRE_WATCH_ISSUER
	claims["aud"] = env.FIRE_WATCH_AUDIENCE
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(duration).Unix()

	tokenString, err := token.SignedString([]byte(env.FIRE_WATCH_JWT_SECRET))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CreateAuthTokens() (string, string, error) {
	env := config.GetCofig()

	accessToken, err := CreateJwtToken(time.Minute * time.Duration(env.FIRE_WATCH_ACCESS_EXPIRED))

	if err != nil {
		return "", "", nil
	}

	refreshToken, err := CreateJwtToken(time.Duration(time.Now().Day()) * time.Duration(env.FIRE_WATCH_REFRESH_EXPIRED))

	if err != nil {
		return "", "", nil
	}

	return accessToken, refreshToken, nil
}
