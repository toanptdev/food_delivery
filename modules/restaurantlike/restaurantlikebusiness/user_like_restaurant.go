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

type IncreaseLikeCountStore interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikeRestaurant struct {
	store         UserLikeStore
	increaseStore IncreaseLikeCountStore
}

func NewUserLikeRestaurant(store UserLikeStore, increaseStore IncreaseLikeCountStore) *userLikeRestaurant {
	return &userLikeRestaurant{store: store, increaseStore: increaseStore}
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

	go func() {
		_ = u.increaseStore.IncreaseLikeCount(ctx, data.RestaurantID)
	}()

	return nil
}
