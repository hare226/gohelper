package rlog

func ShowInConsole(flag bool) {
	printToConsole = flag
}

func ShowInFile(flag bool) {
	printToFile = flag
}

func Debug(msg string) {
	logger.genMsg(Lev_Debug, msg)
	logger.multiWrite()
}

func Info(msg string) {
	logger.genMsg(Lev_Info, msg)
	logger.multiWrite()
}

func Warning(msg string) {
	logger.genMsg(Lev_Warning, msg)
	logger.multiWrite()
}

func Error(msg string) {
	logger.genMsg(Lev_Error, msg)
	logger.multiWrite()
}

// todo 暂时没用
func Panic(msg string) {
	logger.genMsg(Lev_Panic, msg)
	logger.multiWrite()
	panic("rlog调用panic")
}
