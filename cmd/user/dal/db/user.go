package db

import (
	"Demonwuwen/tiktokBackend/pkg/consts"
	"context"
	"errors"
)

type UserLogin struct {
	Id       int64
	UserId   int64
	Username string
	Password string
}

func (l *UserLogin) TableName() string {
	return consts.UserLoginTableName
}

var (
	ErrUsrExist = errors.New("Username already exits ")
)

// CreateUserLogin create userLogin info
func CreateUserLogin(ctx context.Context, userLogin []*UserLogin) error {
	return DB.WithContext(ctx).Create(userLogin).Error
}

func QueryUserLogin(cxt context.Context, username string) ([]*UserLogin, error) {
	res := make([]*UserLogin, 0)
	err := DB.WithContext(cxt).Where("username = ?", username).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
