package jwt

import (
	"time"

	"github.com/Team-Work-Forever/FireWatchRest/config"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
	"github.com/golang-jwt/jwt/v5"
)

type TokenPayload struct {
	UserId   string
	Email    string
	Role     string
	Duration int64
}

func CreateJwtToken(payload TokenPayload) (string, error) {
	env := config.GetCofig()

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = payload.UserId
	claims["email"] = payload.Email
	claims["role"] = payload.Role
	claims["iss"] = env.JWT_ISSUER
	claims["aud"] = env.JWT_AUDIENCE
	claims["iat"] = time.Now().Unix()
	claims["exp"] = payload.Duration

	tokenString, err := token.SignedString([]byte(env.JWT_SECRET))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CreateAuthTokens(auth *entities.Auth) (string, string, error) {
	env := config.GetCofig()

	accessToken, err := CreateJwtToken(TokenPayload{
		Email:    auth.Email.GetValue(),
		UserId:   auth.ID,
		Role:     "admin",
		Duration: time.Now().Add(time.Duration(env.JWT_ACCESS_EXPIRED) * time.Minute).Unix(),
	})

	if err != nil {
		return "", "", nil
	}

	refreshToken, err := CreateJwtToken(TokenPayload{
		Email:    auth.Email.GetValue(),
		UserId:   auth.ID,
		Role:     "admin",
		Duration: time.Now().Add(time.Duration(env.JWT_REFRESH_EXPIRED) * 24 * time.Hour).Unix(),
	})

	if err != nil {
		return "", "", nil
	}

	return accessToken, refreshToken, nil
}
