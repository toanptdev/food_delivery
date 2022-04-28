package userstorage

import (
	"context"
	"rest-api/modules/user/usermodel"
)

func (s *sqlModel) CreateUser(ctx context.Context, user *usermodel.UserCreate) error {
	db := s.db.Begin()

	if err := db.Create(user).Error; err != nil {
		db.Rollback()
		return err
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}

	return nil
}
