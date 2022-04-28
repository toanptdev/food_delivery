package restaurantlikestorage

import (
	"context"
	"rest-api/modules/restaurantlike/restaurantlikemodel"
)

func (s *sqlStore) GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error) {
	result := make(map[int]int)

	db := s.db
	type sqlData struct {
		RestaurantID int `gorm:"column:restaurant_id"`
		LikeCount    int `gorm:"column:like_count"`
	}

	var listLike []sqlData

	err := db.Table(restaurantlikemodel.RestaurantLike{}.TableName()).
		Select("restaurant_id, count(restaurant_id) as like_count").
		Where("restaurant_id in (?)", ids).
		Group("restaurant_id").
		Find(&listLike).Error

	if err != nil {
		return nil, err
	}

	for _, item := range listLike {
		result[item.RestaurantID] = item.LikeCount
	}

	return result, nil
}
