/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2023-08-01 16:35:10
 * @LastEditTime: 2023-08-01 18:08:25
 * @Description: Do not edit
 */
package middleware

import (
	"context"

	"github.com/chenke1115/go-common/configs"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/swag"
)

type SwaggerHandler func(h *server.Hertz, docs *swag.Spec)

// @use h.Use(Swagger(WithSwaggerHandler(h, docs.SwaggerInfo)))
func Swagger(swag SwaggerHandler) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		c.Next(ctx)
	}
}

// swager默认配置处理，可自定义替换
func WithSwaggerHandler(h *server.Hertz, docs *swag.Spec) (f SwaggerHandler) {
	conf := configs.GetConf().Swagger
	if conf.Version != "" {
		docs.Version = conf.Version
	}
	if conf.Host != "" {
		docs.Host = conf.Host
	}
	if conf.BasePath != "" {
		docs.BasePath = conf.BasePath
	}
	if conf.Schemes != nil {
		docs.Schemes = conf.Schemes
	}
	if conf.Title != "" {
		docs.Title = conf.Title
	}

	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))
	return
}
