/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-09-19 13:59:12
 * @LastEditTime: 2023-08-01 17:14:34
 * @Description: Do not edit
 */
package middleware

import (
	"context"
	"fmt"
	"runtime"

	"github.com/chenke1115/hertz-common/pkg/errors"
	"github.com/chenke1115/hertz-common/pkg/response"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

/**
 * @description: Middleware of painc catch
 * @return {*}
 */
func Recovery() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		defer func() {
			// Catch painc error
			if e := recover(); e != nil {
				response.HandleResponse(c, e.(error), nil)
				c.Abort()
			}
		}()

		c.Next(ctx)
	}
}

/**
 * @uses recovery.Recovery(recovery.WithRecoveryHandler(MyRecoveryHandler))
 */
func MyRecoveryHandler(ctx context.Context, c *app.RequestContext, err interface{}, stack []byte) {
	defer func() {
		switch err.(type) {
		case runtime.Error:
			hlog.SystemLogger().CtxErrorf(ctx, "[Recovery] runtime error:%v\nstack:%s", err, stack)
			hlog.SystemLogger().Infof("Client: %s", c.Request.Header.UserAgent())
			c.AbortWithStatus(consts.StatusInternalServerError)
		case string:
			hlog.Errorf(fmt.Sprintf("[Recovery] error:%s", err.(string)))
			response.HandleResponse(c, errors.WrapCode(fmt.Errorf(err.(string)), errors.BadRequest), nil)
			c.AbortWithStatus(consts.StatusBadRequest)
		case error:
			hlog.Errorf(fmt.Sprintf("[Recovery] error:%s", err.(error)))
			response.HandleResponse(c, err.(error), nil)
			c.AbortWithStatus(consts.StatusInternalServerError)
		default:
			hlog.Errorf(fmt.Sprintf("[Recovery] error:%v", err))
			c.AbortWithStatus(consts.StatusInternalServerError)
		}
	}()
}
