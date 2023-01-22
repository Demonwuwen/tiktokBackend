package mw

import (
	tiktokBackendBase "Demonwuwen/tiktokBackend/cmd/api/biz/model/tiktokBackendBase"
	"Demonwuwen/tiktokBackend/pkg/consts"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"time"
)

var JwtMiddleware *jwt.HertzJWTMiddleware

func InitJWT() {
	JwtMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		Key:           []byte(consts.SecretKey),
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		IdentityKey:   consts.IdentityKey,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &tiktokBackendBase.User{
				ID: int64(claims[consts.IdentityKey].(float64)),
			}
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					consts.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var err error
			var req tiktokBackendBase.UserLoginRequest
			if err = c.BindAndValidate(&req); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}
			if len(req.Username) == 0 || len(req.Password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}

			return nil, nil
		},
	})
}
