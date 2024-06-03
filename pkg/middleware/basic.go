/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2023-01-12 11:02:04
 * @LastEditTime: 2023-07-28 09:17:50
 * @Description: Do not edit
 */
package middleware

import (
	"github.com/chenke1115/go-common/configs"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/basic_auth"
)

/**
 * @description: 基础验证
 * @return {*}
 */
func BasicAuth() app.HandlerFunc {
	return basic_auth.BasicAuthForRealm(map[string]string{
		"app_name": configs.GetConf().App.Name,
	}, "", "AppName")
}
