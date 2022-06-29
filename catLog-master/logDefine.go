package catLog

type LogLevel int

const (
	LevelDebug   LogLevel = iota + 1 // 绿色
	LevelInfo                        // 蓝色
	LevelWarning                     // 黄色
	LevelError                       // 红色
)

// outtype
type LogOutType string

const (
	AdapterConsole LogOutType = "console" // 输出到控制台
	AdapterFile    LogOutType = "file"    // 输出到文件
)

type RotatingType int

const (
	TimeRotate RotatingType = iota + 1 // 每隔多长时间回滚
	SizeRotate                         // 固定大小回滚
)

type LogWriter interface {
	// io.Writer

	Write(p []byte) (n int, err error)
}
