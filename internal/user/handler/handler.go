package handler

import (
	"context"
	"github.com/yanguiyuan/cloudspace/internal/cloudfile"
	"github.com/yanguiyuan/cloudspace/internal/user/dal"
	"github.com/yanguiyuan/cloudspace/internal/user/model"
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
	var user *model.User
	err = dal.Q.Transaction(func(tx *dal.Query) error {
		user, err = dal.User.WithContext(ctx).Where(dal.User.Username.Eq(username), dal.User.Password.Eq(password)).FirstOrCreate()
		if err != nil {
			return err
		}
		client, err := cloudfile.NewFileServiceClient()
		if err != nil {
			return err
		}
		r, err := client.CreateFileItem(ctx, username, "namespace", "@root")
		if err != nil {
			return err
		}
		namespaceID, err := client.CreateNamespace(ctx, username, r)
		if err != nil {
			return err
		}
		err = client.CreateUserNamespace(ctx, user.ID, namespaceID, 0)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}
