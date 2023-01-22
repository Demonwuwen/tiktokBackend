// Code generated by hertz generator.

package main

import (
	"Demonwuwen/tiktokBackend/cmd/api/biz/handler"
	"Demonwuwen/tiktokBackend/cmd/api/biz/handler/tiktokBackendBase"
	"Demonwuwen/tiktokBackend/pkg/errno"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)

	// your code ...
	// your code ...
	r.NoRoute(func(ctx context.Context, c *app.RequestContext) { // used for HTTP 404
		tiktokBackendBase.SendResponse(c, errno.ServiceErr, nil)
	})
	r.NoMethod(func(ctx context.Context, c *app.RequestContext) { // used for HTTP 405
		tiktokBackendBase.SendResponse(c, errno.ServiceErr, nil)
	})
}