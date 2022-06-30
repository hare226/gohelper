package main

import (
	"github.com/hare226/gohelper/rlog"
)

func main() {

	rlog.Debug("%s %d %d", "hello", int(3))
	rlog.Info("info")
	rlog.Warning("warnj")
	rlog.Error("error")
	rlog.Panic("panic  %f", "uabi", 0.2)
}
