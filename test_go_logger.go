package main

import (
	"github.com/donnie4w/go-logger/logger"
)

func main() {
	// logger.SetRollingFile("./", "go_log.log", 10, 1, logger.MB)
	// logger.SetLevel(logger.DEBUG)
	// logger.Info("Hello world!")

	//获取一个日志对象
	var log = logger.GetLogger()
	//是否控制台打印
	log.SetConsole(false)
	//按日志大小分割
	// log.SetRollingFile("./", "log.log", 10, 1, logger.MB)
	//按日期分割
	log.SetRollingDaily("./", "log.log")

	log.Info("log hello world")
	log.Error("err hello world")

	//重定向err日志到error.log
	log.SetLevelFile(logger.ERROR, "./", "error.log")
	log.Error("err hello world")
	log.Info("info hello world")
	log.Debug("debug hello world")
}
