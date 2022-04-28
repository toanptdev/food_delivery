package restaurantbusiness

import (
	"context"
	"errors"
	"rest-api/common"

	"rest-api/modules/restaurant/restaurantmodel"
)

type GetRestaurantStore interface {
	GetRestaurantByID(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type getRestaurantBusiness struct {
	store GetRestaurantStore
}

func NewGetRestaurantBusiness(store GetRestaurantStore) *getRestaurantBusiness {
	return &getRestaurantBusiness{store: store}
}

func (g *getRestaurantBusiness) GetRestaurantByID(ctx context.Context, id int) (*restaurantmodel.Restaurant, error) {
	restaurant, err := g.store.GetRestaurantByID(ctx, map[string]interface{}{"id": id})
	if err != nil {
		if err == common.RecordNotFound {
			return nil, common.RecordNotFound
		}

	}
	if restaurant.Status == 0 {
		return nil, errors.New("restaurant has been deleted")
	}
	return restaurant, nil
}
