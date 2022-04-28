package restaurantmodel

import (
	"errors"
	"rest-api/common"
	"rest-api/modules/user/usermodel"
	"strings"
)

type Restaurant struct {
	common.SqlModel
	Name      string          `json:"name,omitempty" form:"name"`
	Addr      string          `json:"addr,omitempty" form:"addr"`
	UserID    int             `json:"-" gorm:"column:owner_id"`
	User      *usermodel.User `json:"user" gorm:"preload:false"`
	LikeCount int             `json:"like_count"`
}

func (r *Restaurant) Validate() error {
	r.Name = strings.TrimSpace(r.Name)
	if r.Name == "" {
		return errors.New("name cant be empty")
	}

	r.Addr = strings.TrimSpace(r.Addr)
	if r.Addr == "" {
		return errors.New("adds cant be empty")
	}

	return nil
}

func (r Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name *string `json:"name" form:"name"`
	Addr *string `json:"addr" form:"addr"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}
