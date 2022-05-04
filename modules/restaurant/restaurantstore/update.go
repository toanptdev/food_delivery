package restaurantstore

import (
	"context"
	"gorm.io/gorm"
	"rest-api/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) UpdateRestaurant(ctx context.Context, id uint, restaurant *restaurantmodel.RestaurantUpdate) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(&restaurant).Error; err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) IncreaseLikeCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.
		Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Update("like_count", gorm.Expr("like_count + ?", 1)).
		Error; err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) DecreaseLikeCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.
		Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Update("like_count", gorm.Expr("like_count - ?", 1)).
		Error; err != nil {
		return err
	}
	return nil
}
