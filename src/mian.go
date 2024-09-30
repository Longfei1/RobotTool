package main

import (
	app2 "RobotTool/src/app"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})

	// 设置 Logrus 的输出为 lumberjack 日志滚动对象
	logger := &lumberjack.Logger{
		Filename:   "logs/app.log", // 日志文件路径
		MaxSize:    10,             // 最大文件大小（MB）
		MaxBackups: 3,              // 最多保留旧文件的数量
		MaxAge:     30,             // 最大文件保存天数
		Compress:   true,           // 是否压缩/归档旧文件
	}
	log.SetOutput(logger)

	app := app2.NewApp()
	app.Run()
}
