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

type DecreaseLikeCountStore interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

type userUnlikeRestaurant struct {
	store         UserUnlikeStore
	decreaseStore DecreaseLikeCountStore
}

func NewUserUnlikeRestaurant(store UserUnlikeStore, decreaseStore DecreaseLikeCountStore) *userUnlikeRestaurant {
	return &userUnlikeRestaurant{store: store, decreaseStore: decreaseStore}
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

	go func() {
		_ = u.decreaseStore.DecreaseLikeCount(ctx, restaurantID)
	}()

	return nil
}
