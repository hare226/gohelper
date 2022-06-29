package main

import (
	"github.com/hare226/gohelper/rlog"
)

func main() {
	var log rlog.LogInf
	// log = loger.NewPrintLog("error")
	log = rlog.NewFileLog(".", "log.log", "error", 17*1024)
	for {
		log.Debug("错误")
		log.Trace("错误")
		log.Info("错误")
		log.Warning("错误")
		log.Error("错误")
		log.Fatal("错误1")
	}
}
