package restaurantbusiness

import (
	"context"
	"rest-api/modules/restaurant/restaurantmodel"
)

type CreateRestaurantStore interface {
	CreateRestaurant(ctx context.Context, restaurant *restaurantmodel.Restaurant) error
}

type createRestaurantBusiness struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBusiness(store CreateRestaurantStore) *createRestaurantBusiness {
	return &createRestaurantBusiness{store: store}
}

func (r *createRestaurantBusiness) CreateRestaurant(ctx context.Context, restaurant *restaurantmodel.Restaurant) error {
	if err := restaurant.Validate(); err != nil {
		return err
	}

	if err := r.store.CreateRestaurant(ctx, restaurant); err != nil {
		return err
	}

	return nil
}
