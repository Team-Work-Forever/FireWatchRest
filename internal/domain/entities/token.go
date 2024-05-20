package entities

import "time"

type Token struct {
	EntityBase
	Token    string    `gorm:"column:token"`
	TType    string    `gorm:"column:type"`
	ExpireAt time.Time `gorm:"column:expire_at"`
}

func NewToken(token string, ttype string, expireAt time.Time) *Token {
	return &Token{
		Token:    token,
		TType:    ttype,
		ExpireAt: expireAt,
	}
}
