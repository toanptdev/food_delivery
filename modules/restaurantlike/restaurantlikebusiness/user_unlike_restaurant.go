package restaurantlikebusiness

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"rest-api/modules/restaurantlike/restaurantlikemodel"
)

type UserUnlikeStore interface {
	Get(ctx context.Context, userID, restaurantID int) (*restaurantlikemodel.RestaurantLike, error)
	Delete(ctx context.Context, userID, restaurantID int) error
}

type userUnlikeRestaurant struct {
	store UserUnlikeStore
}

func NewUserUnlikeRestaurant(store UserUnlikeStore) *userUnlikeRestaurant {
	return &userUnlikeRestaurant{store: store}
}

func (u *userUnlikeRestaurant) Unlike(ctx context.Context, userID, restaurantID int) error {
	_, err := u.store.Get(ctx, userID, restaurantID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("user has not liked this restaurant yet")
		}
		return err
	}

	err = u.store.Delete(ctx, userID, restaurantID)
	if err != nil {
		return err
	}

	return nil
}
