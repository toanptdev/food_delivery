package restaurantbusiness

import (
	"context"
	"rest-api/common"
	"rest-api/modules/restaurant/restaurantmodel"
)

type RestaurantLikeStore interface {
	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
}

type ListRestaurantStore interface {
	ListRestaurantByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]*restaurantmodel.Restaurant, error)
}

type listRestaurantBusiness struct {
	store     ListRestaurantStore
	likeStore RestaurantLikeStore
}

func NewListRestaurantBusiness(store ListRestaurantStore, likeStore RestaurantLikeStore) *listRestaurantBusiness {
	return &listRestaurantBusiness{store: store, likeStore: likeStore}
}

func (r *listRestaurantBusiness) ListRestaurantByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]*restaurantmodel.Restaurant, error) {

	restaurants, err := r.store.ListRestaurantByCondition(ctx, conditions, filter, paging, "User")
	if err != nil {
		return nil, err
	}

	//ids := make([]int, len(restaurants))
	//
	//for i := range restaurants {
	//	ids[i] = restaurants[i].ID
	//}
	//
	//mapResLike, err := r.likeStore.GetRestaurantLikes(ctx, ids)
	//
	//fmt.Println(mapResLike)
	//
	//if mapResLike != nil {
	//	for i, item := range restaurants {
	//		restaurants[i].LikeCount = mapResLike[item.ID]
	//	}
	//}

	return restaurants, nil
}
