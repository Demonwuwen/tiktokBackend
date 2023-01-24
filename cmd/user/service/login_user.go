package service

import (
	"Demonwuwen/tiktokBackend/cmd/user/dal/db"
	"Demonwuwen/tiktokBackend/kitex_gen/user"
	"Demonwuwen/tiktokBackend/pkg/errno"
	"Demonwuwen/tiktokBackend/pkg/util"
	"context"
)

type LoginUserService struct {
	cxt context.Context
}

// NewLoginUserService new NewLoginUserService
func NewLoginUserService(ctx context.Context) *LoginUserService {
	return &LoginUserService{cxt: ctx}
}

func (s *LoginUserService) LoginUser(req *user.UserLoginRequest) (int64, error) {
	username := req.Username
	users, err := db.QueryUserLogin(s.cxt, username)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.AuthorizationFailedErr
	}
	u := users[0]
	ok := util.PasswordVerify(req.Password, u.Password)
	if !ok {
		return 0, errno.AuthorizationFailedErr
	}
	//fmt.Println("u.UserId = ", u.UserId)
	return u.UserId, nil
}
