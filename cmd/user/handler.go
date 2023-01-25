package main

import (
	"Demonwuwen/tiktokBackend/cmd/user/service"
	"Demonwuwen/tiktokBackend/kitex_gen/user"
	"Demonwuwen/tiktokBackend/pkg/consts"
	"Demonwuwen/tiktokBackend/pkg/errno"
	"context"
	"fmt"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	resp, err = service.NewRegisterUserService(ctx).RegisterUser(req)
	return
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	resp = new(user.UserLoginResponse)
	if err = req.IsValid(); err != nil {
		//todo valid resp
		return resp, nil
	}

	uid, err := service.NewLoginUserService(ctx).LoginUser(req)
	if err != nil {
		return resp, err
	}

	resp.UserId = uid
	resp.StatusCode = consts.SuccessCode
	resp.StatusMsg = errno.Success.ErrMsg
	//fmt.Println("Uid = ", uid)
	return resp, nil
}

// UserGet implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserGet(ctx context.Context, req *user.UserRequest) (resp *user.UserResponse, err error) {
	// TODO: Your code here...
	resp = &user.UserResponse{}
	resp, err = service.NewGetUserService(ctx).GetUserById(req.UserId)
	if err != nil {
		fmt.Println("UserGet err :", err)
	}
	return resp, err
}
