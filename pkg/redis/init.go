/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-30 16:12:25
 * @LastEditTime: 2023-01-13 15:23:42
 * @Description: Do not edit
 */
package redis

import (
	"context"
	"time"

	"github.com/chenke1115/go-common/configs"
	"github.com/chenke1115/hertz-common/global"
	"github.com/go-redis/redis/v8" // notice: new version v8
)

/**
 * @description: Init redis client
 * @return {*}
 */
func InitClient(conf *configs.Options) (err error) {
	rConf := conf.Redis
	global.RedisDB = redis.NewClient(&redis.Options{
		Addr:     rConf.Addr,
		Password: rConf.Password,
		DB:       rConf.DB,
		PoolSize: rConf.Size,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = global.RedisDB.Ping(ctx).Result()
	return err
}
