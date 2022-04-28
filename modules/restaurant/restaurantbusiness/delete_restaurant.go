package restaurantbusiness

import (
	"context"
	"errors"

	"rest-api/modules/restaurant/restaurantmodel"
)

type DeleteRestaurantStore interface {
	GetRestaurantByID(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
	DeleteRestaurant(ctx context.Context, id uint) error
}

type deleteRestaurantBusiness struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBusiness(store DeleteRestaurantStore) *deleteRestaurantBusiness {
	return &deleteRestaurantBusiness{store: store}
}

func (r *deleteRestaurantBusiness) DeleteRestaurant(ctx context.Context, id uint) error {
	oldRestaurant, err := r.store.GetRestaurantByID(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if oldRestaurant.Status == 0 {
		return errors.New("restaurant has been deleted")
	}
	if err := r.store.DeleteRestaurant(ctx, id); err != nil {
		return err
	}

	return nil
}
