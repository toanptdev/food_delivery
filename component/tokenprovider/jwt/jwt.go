package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"rest-api/component/tokenprovider"
	"time"
)

type jwtTokenProvider struct {
	SecretKey string
}

type MyClaims struct {
	Payload tokenprovider.Payload
	jwt.StandardClaims
}

func NewJwtTokenProvider(secretKey string) *jwtTokenProvider {
	return &jwtTokenProvider{SecretKey: secretKey}
}

func (j *jwtTokenProvider) Generate(payload tokenprovider.Payload, expiry int) (*tokenprovider.Token, error) {
	data := MyClaims{
		Payload: tokenprovider.Payload{
			UserID: payload.UserID,
			Role:   payload.Role,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(expiry) * time.Second).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	ss, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return nil, err
	}
	return &tokenprovider.Token{
		Token:   ss,
		Expiry:  expiry,
		Created: time.Now(),
	}, nil
}

func (j *jwtTokenProvider) Validate(token string) (*tokenprovider.Payload, error) {
	ss, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := ss.Claims.(*MyClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	return &claims.Payload, nil
}
