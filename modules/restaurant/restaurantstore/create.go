package restaurantstore

import (
	"context"
	"rest-api/common"

	"rest-api/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) CreateRestaurant(ctx context.Context, restaurant *restaurantmodel.Restaurant) error {
	db := s.db
	if err := db.Create(restaurant).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
