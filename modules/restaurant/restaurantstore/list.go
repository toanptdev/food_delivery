package restaurantstore

import (
	"context"
	"rest-api/common"

	"rest-api/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) ListRestaurantByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]*restaurantmodel.Restaurant, error) {
	var restaurants []*restaurantmodel.Restaurant
	db := s.db

	db = db.Table(restaurantmodel.Restaurant{}.TableName()).Where(conditions).Where("status IN (1)")
	if filter != nil {
		if filter.Name != "" {
			db = db.Where("name = ?", filter.Name)
		}

		if filter.Addr != "" {
			db = db.Where("name = ?", filter.Addr)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for k := range moreKeys {
		db = db.Preload(moreKeys[k])
	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&restaurants).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return restaurants, nil
}
