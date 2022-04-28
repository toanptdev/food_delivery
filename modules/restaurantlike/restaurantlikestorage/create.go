package restaurantlikestorage

import (
	"context"
	"rest-api/modules/restaurantlike/restaurantlikemodel"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantlikemodel.RestaurantLike) error {
	db := s.db.Begin()

	if err := db.Create(&data).Error; err != nil {
		db.Rollback()
		return err
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}

	return nil
}
