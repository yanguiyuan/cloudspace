package handler

import (
	"context"
	"github.com/yanguiyuan/cloudspace/internal/user/dal"
	"github.com/yanguiyuan/cloudspace/internal/user/model"
	"github.com/yanguiyuan/yuan/pkg/gen/id"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, username string, password string) (resp int64, err error) {
	// TODO: Your code here...
	find, err := dal.User.WithContext(ctx).Where(dal.User.Username.Eq(username), dal.User.Password.Eq(password)).First()
	if err != nil {
		return 0, err
	}
	return find.ID, nil
}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, username string, password string) (resp int64, err error) {
	i := int64(id.One())
	err = dal.User.WithContext(ctx).Create(&model.User{ID: i, Username: username, Password: password})
	if err != nil {
		return 0, err
	}
	return i, nil
}
