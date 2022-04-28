package userstorage

import (
	"context"
	"rest-api/modules/user/usermodel"
)

func (s *sqlModel) GetUser(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*usermodel.User, error) {
	db := s.db
	var user usermodel.User

	for k := range moreKeys {
		db.Preload(moreKeys[k])
	}

	if err := db.Where(conditions).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
