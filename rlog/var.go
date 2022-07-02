package rlog

// LogLevel 自定义类型
type LogLevel string

// 日志级别常量
const (
	Lev_Debug   LogLevel = "[DEBUG]"
	Lev_Info    LogLevel = "[ INFO]"
	Lev_Warning LogLevel = "[ WARN]"
	Lev_Error   LogLevel = "[ERROR]"
	Lev_Panic   LogLevel = "[PANIC]"
)

const (
	DateFormat      = `%s-%02d-%02d`
	LogNameFormat   = `%s_%s.r.log`
	LogRelativePath = "log"

	TimeFormat  = `%02d:%02d:%02d:%03d`
	CallerDepth = 2
	MsgFromat   = `%s %s [%s:%d] %s`
	Separator   = "/"
)

var (
	printToConsole = true
	printToFile    = true
)

var logger *LogEntity

func init() {
	logger = new(LogEntity)
	logger.file = genLogFile()
}
