package main

import (
	"context"
	tiktokbackendbase "github.com/demonwuwen/tiktokBackend/kitex_gen/tiktokBackendBase"
)

// BaseApiServiceImpl implements the last service interface defined in the IDL.
type BaseApiServiceImpl struct{}

// UserRegister implements the BaseApiServiceImpl interface.
func (s *BaseApiServiceImpl) UserRegister(ctx context.Context, req *tiktokbackendbase.UserRegisterRequest) (resp *tiktokbackendbase.UserRegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// UserLogin implements the BaseApiServiceImpl interface.
func (s *BaseApiServiceImpl) UserLogin(ctx context.Context, req *tiktokbackendbase.UserLoginRequest) (resp *tiktokbackendbase.UserLoginResponse, err error) {
	// TODO: Your code here...

	return
}

// UserGet implements the BaseApiServiceImpl interface.
func (s *BaseApiServiceImpl) UserGet(ctx context.Context, req *tiktokbackendbase.UserRequest) (resp *tiktokbackendbase.UserResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishVideo implements the BaseApiServiceImpl interface.
func (s *BaseApiServiceImpl) PublishVideo(ctx context.Context, req *tiktokbackendbase.PublishActionRequest) (resp *tiktokbackendbase.PublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishList implements the BaseApiServiceImpl interface.
func (s *BaseApiServiceImpl) PublishList(ctx context.Context, req *tiktokbackendbase.PublishListRequest) (resp *tiktokbackendbase.PublishListResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteNote implements the BaseApiServiceImpl interface.
func (s *BaseApiServiceImpl) DeleteNote(ctx context.Context, req *tiktokbackendbase.FeedRequest) (resp *tiktokbackendbase.FeedResponse, err error) {
	// TODO: Your code here...
	return
}
