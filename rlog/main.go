package rlog

import (
	"fmt"
)

func init() {
	logger = new(LogEntity)
	logger.file = genLogFile()
}

func ShowInConsole(flag bool) {
	printToConsole = flag
}

func ShowInFile(flag bool) {
	printToFile = flag
}

func Debug(msg string, a ...any) {
	baseMsg := fmt.Sprintf(msg, a)
	logger.genMsg(Lev_Debug, baseMsg)
	logger.multiWrite()
}

func Info(msg string, a ...any) {
	baseMsg := fmt.Sprintf(msg, a)
	logger.genMsg(Lev_Info, baseMsg)
	logger.multiWrite()
}

func Error(msg string, a ...any) {
	var con string
	if len(a) > 0 {
		con = fmt.Sprintf(msg, a)
	} else {
		con = fmt.Sprintf(msg)
	}

	logger.genMsg(Lev_Error, con)
	logger.multiWrite()
}

func Warning(msg string, a ...any) {
	baseMsg := fmt.Sprintf(msg, a)
	logger.genMsg(Lev_Warning, baseMsg)
	logger.multiWrite()
}

// todo 暂时没用
func Panic(msg string, a ...any) {
	baseMsg := fmt.Sprintf(msg, a)
	logger.genMsg(Lev_Panic, baseMsg)
	logger.multiWrite()
}
