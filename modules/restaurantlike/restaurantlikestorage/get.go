package restaurantlikestorage

import (
	"context"
	"rest-api/modules/restaurantlike/restaurantlikemodel"
)

func (s *sqlStore) Get(ctx context.Context, userID, restaurantID int) (*restaurantlikemodel.RestaurantLike, error) {
	var userLike *restaurantlikemodel.RestaurantLike
	db := s.db
	if err := db.Where("user_id = ? and restaurant_id = ?", userID, restaurantID).First(&userLike).Error; err != nil {
		return nil, err
	}

	return userLike, nil
}
