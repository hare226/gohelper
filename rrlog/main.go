package main

import (
	"time"

	"github.com/hare226/gohelper/rlog/mylogger"
)

// 测试我们自己写的日志库
func main() {
	log := mylogger.NewLog("info")
	for {
		log.Debug("这是一条Degub日志")
		log.Info("这是一条info日志")
		log.Warning("这是一条Warning日志")
		log.Error("这是一条Error日志")
		log.Fatal("这是一条Fatal日志")
		time.Sleep(time.Second)
	}
}
