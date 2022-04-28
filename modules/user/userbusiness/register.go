package userbusiness

import (
	"context"
	"errors"
	"fmt"
	"rest-api/common"
	"rest-api/component/hasher"
	"rest-api/modules/user/usermodel"
)

type RegisterStorage interface {
	GetUser(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, user *usermodel.UserCreate) error
}

type userBusiness struct {
	store RegisterStorage
	hash  *hasher.MD5Hash
}

func NewUserBusiness(store RegisterStorage, hash *hasher.MD5Hash) *userBusiness {
	return &userBusiness{store: store, hash: hash}
}

func (u *userBusiness) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := u.store.GetUser(ctx, map[string]interface{}{"email": data.Email})
	if user != nil {
		return errors.New("email already used")
	}

	salt := common.GenSalt(50)
	hashedPassword := u.hash.Hash(data.Password + salt)
	data.Salt = salt
	fmt.Println("salt", salt)
	fmt.Println("input", data.Password)
	data.Password = hashedPassword
	data.Role = "user"
	if err := u.store.CreateUser(ctx, data); err != nil {
		return err
	}

	return nil
}
