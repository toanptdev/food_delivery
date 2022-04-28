package restaurantstore

import (
	"context"
	"rest-api/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) UpdateRestaurant(ctx context.Context, id uint, restaurant *restaurantmodel.RestaurantUpdate) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(&restaurant).Error; err != nil {
		return err
	}
	return nil
}
