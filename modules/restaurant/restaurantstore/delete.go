package restaurantstore

import (
	"context"
	"rest-api/common"
	"rest-api/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) DeleteRestaurant(ctx context.Context, id uint) error {
	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
