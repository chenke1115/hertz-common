/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2023-06-08 15:29:22
 * @LastEditTime: 2023-06-08 15:42:04
 * @Description: Do not edit
 */
package middleware

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/chenke1115/hertz-common/pkg/errors"
	"github.com/chenke1115/hertz-common/pkg/response"
	"github.com/cloudwego/hertz/pkg/app"
)

// 令牌桶结构
type TokenBucket struct {
	capacity  int64      // 桶的容量
	rate      float64    // 令牌放入速率
	tokens    float64    // 当前令牌数量
	lastToken time.Time  // 上一次放令牌的时间
	mtx       sync.Mutex // 互斥锁
}

/**
 * @description: 限流逻辑
 * @return {*}
 */
func (tb *TokenBucket) Allow() bool {
	tb.mtx.Lock()
	defer tb.mtx.Unlock()

	now := time.Now()
	// 计算需要放的令牌数量
	tb.tokens = tb.tokens + tb.rate*now.Sub(tb.lastToken).Seconds()
	if tb.tokens > float64(tb.capacity) {
		tb.tokens = float64(tb.capacity)
	}
	// 判断是否允许请求
	if tb.tokens >= 1 {
		tb.tokens--
		tb.lastToken = now
		return true
	} else {
		return false
	}
}

/**
 * @description: 限流中间件
 * @param {int64} maxConn  最大链接数
 * @return {*}
 */
func Limithander(maxConn int64) app.HandlerFunc {
	tb := &TokenBucket{
		capacity:  maxConn,
		rate:      1.0,
		tokens:    0,
		lastToken: time.Now(),
	}

	return func(c context.Context, ctx *app.RequestContext) {
		if !tb.Allow() {
			response.HandleResponse(ctx, errors.WrapCode(fmt.Errorf("Too many request"), 503), nil)
			ctx.Abort()
			return
		}
		ctx.Next(c)
	}
}
