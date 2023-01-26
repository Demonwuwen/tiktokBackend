package rpc

import (
	"Demonwuwen/tiktokBackend/kitex_gen/user"
	"Demonwuwen/tiktokBackend/kitex_gen/user/userservice"
	"Demonwuwen/tiktokBackend/pkg/consts"
	"Demonwuwen/tiktokBackend/pkg/errno"
	"Demonwuwen/tiktokBackend/pkg/mw"
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func initUser() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ApiServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := userservice.NewClient(
		consts.UserServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

func UserLogin(ctx context.Context, req *user.UserLoginRequest) (int64, error) {
	resp, err := userClient.UserLogin(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 0 {
		msg := resp.StatusMsg
		return 0, errno.NewErrNo(resp.StatusCode, msg)
	}

	return resp.UserId, nil
}

func RegisterUser(ctx context.Context, req *user.UserRegisterRequest) (*user.UserRegisterResponse, error) {
	resp, err := userClient.UserRegister(ctx, req)
	if err != nil {
		fmt.Println("register failed ")
		return resp, err
	}

	fmt.Println("rpc Register")
	return resp, nil
}

func GetUser(ctx context.Context, req *user.UserRequest) (*user.UserResponse, error) {
	resp, err := userClient.UserGet(ctx, req)
	if err != nil {
		fmt.Println("get user failed ")

		return resp, err
	}

	fmt.Println("get user success")
	return resp, nil
}
