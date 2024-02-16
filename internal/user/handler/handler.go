package handler

import (
	"context"
	"github.com/yanguiyuan/cloudspace/internal/cloudfile"
	"github.com/yanguiyuan/cloudspace/internal/user/dal"
	"github.com/yanguiyuan/cloudspace/internal/user/model"
	"github.com/yanguiyuan/cloudspace/pkg/rpc"
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
		namespaceID, err := client.CreateNamespace(ctx, username)
		if err != nil {
			return err
		}
		_, err = client.CreateFileItem(ctx, username, "namespace", "@root", namespaceID)
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
func (s *UserServiceImpl) GetUser(ctx context.Context, id int64) (resp *rpc.User, err error) {
	// TODO: Your code here...
	user, err := dal.User.WithContext(ctx).Where(dal.User.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return &rpc.User{
		Id:         user.ID,
		Username:   user.Username,
		Password:   user.Password,
		Gender:     user.Gender,
		Email:      user.Email,
		Role:       user.Role,
		CreateTime: user.CreateTime.String(),
		UpdateTime: user.UpdateTime.String(),
	}, nil
}
func (s *UserServiceImpl) UpdateUser(ctx context.Context, user *rpc.User) (err error) {
	_, err = dal.User.WithContext(ctx).Where(dal.User.ID.Eq(user.Id)).Updates(map[string]interface{}{"gender": user.Gender, "email": user.Email})
	return err
}
func (s *UserServiceImpl) GetUsers(ctx context.Context, offset int32, limit int32) (resp []*rpc.User, err error) {
	var res []*rpc.User
	r, err := dal.User.WithContext(ctx).Limit(int(limit)).Offset(int(offset)).Find()
	if err != nil {
		return nil, err
	}
	for _, v := range r {
		res = append(res, &rpc.User{
			Id:         v.ID,
			Username:   v.Username,
			Password:   v.Password,
			Gender:     v.Gender,
			Email:      v.Email,
			Role:       v.Role,
			CreateTime: v.CreateTime.String(),
			UpdateTime: v.UpdateTime.String(),
		})
	}
	return res, nil
}
