package catLog

import (
	"time"
)

type logMsg struct {
	level LogLevel
	msg   string
	when  time.Time
}
