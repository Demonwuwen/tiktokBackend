package service

import (
	"Demonwuwen/tiktokBackend/cmd/user/dal/db"
	"Demonwuwen/tiktokBackend/kitex_gen/user"
	"Demonwuwen/tiktokBackend/pkg/errno"
	"Demonwuwen/tiktokBackend/pkg/util"
	"context"
	"fmt"
)

type RegisterUserService struct {
	ctx context.Context
}

func NewRegisterUserService(ctx context.Context) *RegisterUserService {
	return &RegisterUserService{ctx: ctx}
}

func (s *RegisterUserService) RegisterUser(req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	users, err := db.QueryUserLogin(s.ctx, req.Username)
	resp = &user.UserRegisterResponse{}

	if err != nil {
		return resp, err
	}

	if len(users) != 0 {
		resp.StatusCode = errno.UserAlreadyExistErr.ErrCode
		resp.StatusMsg = errno.UserAlreadyExistErr.ErrMsg
		return resp, nil
	}

	user := db.User{}
	id, err := db.CreateUser(context.Background(), &user)
	fmt.Println("createUser id = ", id)
	h, _ := util.GenPasswordHash(req.Password)

	err = db.CreateUserLogin(s.ctx, []*db.UserLogin{{
		Username: req.Username,
		UserId:   user.Id,
		Password: h,
	}})
	if err != nil {
		return resp, err
	}
	resp.UserId = id
	resp.StatusMsg = errno.Success.ErrMsg
	return resp, nil

}
