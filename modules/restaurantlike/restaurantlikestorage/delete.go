package restaurantlikestorage

import (
	"context"
	"rest-api/modules/restaurantlike/restaurantlikemodel"
)

func (s *sqlStore) Delete(ctx context.Context, userID, restaurantID int) error {
	db := s.db
	if err := db.
		Table(restaurantlikemodel.RestaurantLike{}.TableName()).
		Where("user_id = ? and restaurant_id = ?", userID, restaurantID).
		Delete(nil).Error; err != nil {
		return err
	}

	return nil
}
