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

type RotatingType int

type LogWriter interface {
	Write(p []byte) (n int, err error)
}
