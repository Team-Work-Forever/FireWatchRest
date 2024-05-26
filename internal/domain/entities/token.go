package entities

import (
	"time"

	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/key"
)

type (
	TokenType struct {
		Key string
		Exp time.Duration
	}

	Token struct {
		Token    string
		Type     string
		ExpireAt time.Duration
	}
)

var (
	ForgotToken TokenType = TokenType{
		Key: "forgot-token",
		Exp: time.Minute * 15,
	}
)

func NewToken(
	token string,
	tt TokenType,
) *Token {
	return &Token{
		Token:    token,
		Type:     tt.GetType(),
		ExpireAt: tt.Exp,
	}
}

func (td *TokenType) GetType() string {
	return td.Key
}

func (t *Token) GetKV() key.KeyValue {
	return key.KeyValue{
		Key: key.Key{
			Tag:   t.Type,
			Value: t.Token,
		},
		Value: t.Token,
		Exp:   t.ExpireAt,
	}
}

func (t *Token) Scan(key key.Key, value string) error {
	t.Token = value
	t.Type = key.Tag

	return nil
}
