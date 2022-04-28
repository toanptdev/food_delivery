package restaurantbusiness

import (
	"context"
	"errors"

	"rest-api/modules/restaurant/restaurantmodel"
)

type UpdateRestaurantStore interface {
	GetRestaurantByID(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
	UpdateRestaurant(ctx context.Context, id uint, restaurant *restaurantmodel.RestaurantUpdate) error
}

type updateRestaurantBusiness struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBusiness(store UpdateRestaurantStore) *updateRestaurantBusiness {
	return &updateRestaurantBusiness{store: store}
}

func (r *updateRestaurantBusiness) UpdateRestaurant(ctx context.Context, id uint, restaurant *restaurantmodel.RestaurantUpdate) error {
	oldRestaurant, err := r.store.GetRestaurantByID(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if oldRestaurant.Status == 0 {
		return errors.New("restaurant has been deleted")
	}
	if err := r.store.UpdateRestaurant(ctx, id, restaurant); err != nil {
		return err
	}

	return nil
}
