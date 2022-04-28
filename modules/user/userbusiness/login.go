package userbusiness

import (
	"context"
	"errors"
	"fmt"
	"rest-api/component/hasher"
	"rest-api/component/tokenprovider"
	"rest-api/modules/user/usermodel"
)

type LoginStorage interface {
	GetUser(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*usermodel.User, error)
}

type loginBusiness struct {
	store         LoginStorage
	hash          *hasher.MD5Hash
	tokenProvider tokenprovider.Provider
	Expiry        int
}

func NewLoginBusiness(store LoginStorage, hash *hasher.MD5Hash, tokenProvider tokenprovider.Provider, expiry int) *loginBusiness {
	return &loginBusiness{
		store:         store,
		hash:          hash,
		tokenProvider: tokenProvider,
		Expiry:        expiry,
	}
}

func (l *loginBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	user, err := l.store.GetUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, err
	}

	hashedPassword := l.hash.Hash(data.Password + user.Salt)
	if user.Password != hashedPassword {
		fmt.Println("input", data.Password)
		fmt.Println("salted", user.Salt)
		fmt.Println("hashed ", hashedPassword)
		fmt.Println("pass ", user.Password)
		return nil, errors.New("email or password is invalid")
	}

	payload := tokenprovider.Payload{
		UserID: user.ID,
		Role:   user.Role,
	}

	fmt.Println(l.Expiry)

	token, err := l.tokenProvider.Generate(payload, l.Expiry)
	if err != nil {
		return nil, err
	}

	return token, nil
}
