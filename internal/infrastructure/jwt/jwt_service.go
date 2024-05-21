package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/Team-Work-Forever/FireWatchRest/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type (
	AuthTokenPayload struct {
		Email  string
		UserId string
		Role   string
	}

	TokenPayload struct {
		UserId   string
		Email    string
		Role     string
		Duration time.Time
	}

	AuthClaims struct {
		jwt.RegisteredClaims
		Email string `json:"email,omitempty"`
		Role  string `json:"role,omitempty"`
	}
)

func CreateJwtToken(payload TokenPayload) (string, error) {
	env := config.GetCofig()

	claims := AuthClaims{
		Email: payload.Email,
		Role:  payload.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    env.JWT_ISSUER,
			Audience:  jwt.ClaimStrings{env.JWT_AUDIENCE},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
			ExpiresAt: &jwt.NumericDate{Time: payload.Duration},
			Subject:   payload.UserId,
			ID:        uuid.New().String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(env.JWT_SECRET))
}

func CreateAuthTokens(payload AuthTokenPayload) (string, string, error) {
	env := config.GetCofig()

	accessToken, err := CreateJwtToken(TokenPayload{
		Email:    payload.Email,
		UserId:   payload.UserId,
		Role:     payload.Role,
		Duration: time.Now().Add(time.Duration(env.JWT_ACCESS_EXPIRED) * time.Minute),
	})

	if err != nil {
		return "", "", nil
	}

	refreshToken, err := CreateJwtToken(TokenPayload{
		Email:    payload.Email,
		UserId:   payload.UserId,
		Role:     payload.Role,
		Duration: time.Now().Add(time.Duration(env.JWT_REFRESH_EXPIRED) * 24 * time.Hour),
	})

	if err != nil {
		return "", "", nil
	}

	return accessToken, refreshToken, nil
}

func GetClaims(token string, claims jwt.Claims) (interface{}, error) {
	env := config.GetCofig()

	jwtToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(env.JWT_SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	if !jwtToken.Valid {
		return nil, errors.New("token is no longer valid")
	}

	return claims, nil
}

func ValidateToken(token string) bool {
	_, err := GetClaims(token, &AuthClaims{})

	return err == nil
}
