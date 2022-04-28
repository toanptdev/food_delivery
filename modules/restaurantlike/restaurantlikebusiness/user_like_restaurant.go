package restaurantlikebusiness

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"rest-api/modules/restaurantlike/restaurantlikemodel"
)

type UserLikeStore interface {
	Get(ctx context.Context, userID, restaurantID int) (*restaurantlikemodel.RestaurantLike, error)
	Create(ctx context.Context, data *restaurantlikemodel.RestaurantLike) error
}

type userLikeRestaurant struct {
	store UserLikeStore
}

func NewUserLikeRestaurant(store UserLikeStore) *userLikeRestaurant {
	return &userLikeRestaurant{store: store}
}

func (u *userLikeRestaurant) Like(ctx context.Context, data *restaurantlikemodel.RestaurantLike) error {
	liked, err := u.store.Get(ctx, data.UserID, data.RestaurantID)
	if err != gorm.ErrRecordNotFound {
		return err
	}
	if liked != nil {
		return errors.New("user had liked this restaurant")
	}

	err = u.store.Create(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
