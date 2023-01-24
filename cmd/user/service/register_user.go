package service

import (
	"Demonwuwen/tiktokBackend/cmd/user/dal/db"
	"Demonwuwen/tiktokBackend/kitex_gen/user"
	"Demonwuwen/tiktokBackend/pkg/errno"
	"Demonwuwen/tiktokBackend/pkg/util"
	"context"
)

type RegisterUserService struct {
	ctx context.Context
}

func (s *RegisterUserService) RegisterUser(req *user.UserRegisterRequest) error {
	users, err := db.QueryUserLogin(s.ctx, req.Username)
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}
	h, _ := util.GenPasswordHash(req.Password)
	return db.CreateUserLogin(s.ctx, []*db.UserLogin{{
		Username: req.Username,
		Password: h,
	}})

}
