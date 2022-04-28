package restaurantlikemodel

import "time"

type RestaurantLike struct {
	RestaurantID int       `json:"restaurant_id" gorm:"column:restaurant_id"`
	UserID       int       `json:"user_id" gorm:"column:user_id"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
}

func (r RestaurantLike) TableName() string {
	return "restaurant_likes"
}
