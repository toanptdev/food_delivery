package usermodel

import "rest-api/common"

type User struct {
	common.SqlModel
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"-" gorm:"column:password"`
	Salt     string `json:"-" gorm:"column:salt"`
	Role     string `json:"role" gorm:"column:role"`
}

func (u User) TableName() string {
	return "users"
}

type UserCreate struct {
	common.SqlModel
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
	Salt     string `json:"-" gorm:"column:salt"`
	Role     string `json:"-" gorm:"column:role"`
}

func (u UserCreate) TableName() string {
	return User{}.TableName()
}

type UserLogin struct {
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}
