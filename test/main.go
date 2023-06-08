/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2023-06-08 16:53:31
 * @LastEditTime: 2023-06-08 17:32:38
 * @Description: Do not edit
 */
package main

import (
	"context"
	"time"

	"github.com/chenke1115/hertz-common/pkg/middleware"
	"github.com/chenke1115/hertz-common/pkg/response"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func main() {
	// server.Default() creates a Hertz with recovery middleware.
	// Maximum wait time before exit, if not specified the default is 5s
	h := server.Default(
		server.WithExitWaitTime(0 * time.Second),
	)

	h.Use(middleware.Limithander(100))

	h.GET("/", func(c context.Context, ctx *app.RequestContext) {
		response.HandleResponse(ctx, nil, "Hello World!")
	})

	// Graceful exit
	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		<-ctx.Done()
		hlog.Warn("exit timeout!")
	})

	// run
	h.Spin()
}
