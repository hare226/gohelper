package rlog

// LogLevel 自定义类型
type LogLevel string

// 日志级别常量
const (
	Lev_Debug   LogLevel = "[DEBUG]"
	Lev_Info    LogLevel = "[INFO]"
	Lev_Warning LogLevel = "[WARNING]"
	Lev_Error   LogLevel = "[ERROR]"
	Lev_Panic   LogLevel = "[PANIC]"
)

const (
	DateFormat      = `%s-%02d-%02d`
	TimeFormat      = `%02d:%02d:%02d:%03d`
	LogNameFormat   = `%s_%s.r.log`
	MsgFromat       = `%s %s [%s:%d] %s`
	LogRelativePath = "log"
	CallerDepth     = 2
)

var (
	printToConsole = true
	printToFile    = true
)

var logger *LogEntity
