package service

import (
	"Demonwuwen/tiktokBackend/cmd/user/dal/db"
	"Demonwuwen/tiktokBackend/kitex_gen/user"
	"Demonwuwen/tiktokBackend/pkg/errno"
	"context"
)

type GetUserService struct {
	cxt context.Context
}

// NewGetUserService  new NewGetUserService
func NewGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{cxt: ctx}
}

func (s *GetUserService) GetUserById(userId int64) (*user.UserResponse, error) {
	resp := &user.UserResponse{}
	user, err := db.QueryUserById(s.cxt, userId)
	if err != nil {
		return resp, err
	}
	if user.Id == 0 {
		resp.StatusCode = errno.UserNotFoundErr.ErrCode
		resp.StatusMsg = errno.UserAlreadyExistErr.ErrMsg
		return resp, errno.AuthorizationFailedErr
	}
	resp.StatusCode = errno.Success.ErrCode
	resp.StatusMsg = errno.Success.ErrMsg
	resp.User = user
	//fmt.Println("u.UserId = ", u.UserId)
	return resp, nil
}
