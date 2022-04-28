package tokenprovider

import "time"

type Provider interface {
	Generate(payload Payload, expiry int) (*Token, error)
	Validate(token string) (*Payload, error)
}

type Payload struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
}

type Token struct {
	Token   string    `json:"token"`
	Expiry  int       `json:"expiry"`
	Created time.Time `json:"created"`
}
