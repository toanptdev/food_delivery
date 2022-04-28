package restaurantstore

import (
	"context"
	"gorm.io/gorm"
	"rest-api/common"
	"rest-api/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) GetRestaurantByID(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	db := s.db
	for k := range moreKeys {
		db = db.Preload(moreKeys[k])
	}
	var restaurant restaurantmodel.Restaurant
	if err := db.Where(conditions).First(&restaurant).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &restaurant, nil
}
