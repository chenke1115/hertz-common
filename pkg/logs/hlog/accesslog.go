/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2023-01-03 17:54:03
 * @LastEditTime: 2023-01-04 11:45:20
 * @Description: Do not edit
 */
package hlog

import (
	"fmt"
	"os"
	"time"

	"github.com/chenke1115/go-common/configs"
	"github.com/chenke1115/go-common/functions/file"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzzap "github.com/hertz-contrib/logger/zap"
	"gopkg.in/natefinch/lumberjack.v2"
)

/**
 * @description: write log to logfile
 * @return {*}
 */
func WriteLog(conf *configs.Options) {
	path := conf.Log.Dir + "hlog/"
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

	// For zap detailed settings
	logger := hertzzap.NewLogger()
	// Provides compression and deletion
	lumberjackLogger := &lumberjack.Logger{
		Filename:   path + fileName,
		MaxSize:    conf.Log.MaxSize,
		MaxBackups: conf.Log.MaxBackups,
		MaxAge:     conf.Log.MaxAge,
		Compress:   conf.Log.Compress,
	}

	// set log level
	level := hlog.LevelInfo
	if conf.Debug {
		level = hlog.LevelDebug
	}

	logger.SetOutput(lumberjackLogger)
	logger.SetLevel(level)
	hlog.SetLogger(logger)
}
