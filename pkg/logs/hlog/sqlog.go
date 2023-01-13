/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2023-01-03 17:54:27
 * @LastEditTime: 2023-01-04 11:44:53
 * @Description: Do not edit
 */
package hlog

import (
	"fmt"
	"os"
	"time"

	"github.com/chenke1115/go-common/configs"
	"github.com/chenke1115/go-common/functions/file"
	"gopkg.in/natefinch/lumberjack.v2"
)

/**
 * @description: log of sql
 * @param {*configs.Options} conf
 * @return {*}
 */
func SqlLog(conf *configs.Options) *lumberjack.Logger {
	path := conf.Log.Dir + "sqlog/"
	fileName := time.Now().Format("20060102") + ".log"
	_, err := os.Stat(path + fileName)
	switch {
	case os.IsNotExist(err):
		if err = file.MakeDir(path); err != nil {
			panic(fmt.Errorf("log file not exist:%v", err.Error()))
		}
	case os.IsPermission(err):
		panic(fmt.Errorf("log file permission:%v", err.Error()))
	}

	// Provides compression and deletion
	lumberjackLogger := &lumberjack.Logger{
		Filename:   path + fileName,
		MaxSize:    conf.Log.MaxSize,
		MaxBackups: conf.Log.MaxBackups,
		MaxAge:     conf.Log.MaxAge,
		Compress:   conf.Log.Compress,
	}

	return lumberjackLogger
}
