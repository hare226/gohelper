package mylogger

import (
	"fmt"
	"time"
)

// 日志级别
type LogLevel uint16

// Logger 日志结构体
type Logger struct {
	Level LogLevel
}

// 向终端写日志相关内容
// NewLog 构造函数
func NewLog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		Level: level,
	}
}
func (l Logger) Debug(msg string) {
	if l.Level <= DEBUG {
		now := time.Now()
		funcName, fileName, lineNo := getInfo(2)
		fmt.Printf("[%s][DEBUG] [%s: %s: %d]%s\n", now.Format("2006-01-02 15:04:05"), funcName, fileName, lineNo, msg)
	}
}
func (l Logger) Info(msg string) {
	if l.Level <= INFO {
		now := time.Now()
		funcName, fileName, lineNo := getInfo(2)
		fmt.Printf("[%s][DEBUG] [%s: %s: %d]%s\n", now.Format("2006-01-02 15:04:05"), funcName, fileName, lineNo, msg)
	}
}
func (l Logger) Warning(msg string) {
	if l.Level <= WARNING {
		now := time.Now()
		funcName, fileName, lineNo := getInfo(2)
		fmt.Printf("[%s][DEBUG] [%s: %s: %d]%s\n", now.Format("2006-01-02 15:04:05"), funcName, fileName, lineNo, msg)
	}
}
func (l Logger) Error(msg string) {
	if l.Level <= ERROR {
		now := time.Now()
		funcName, fileName, lineNo := getInfo(2)
		fmt.Printf("[%s][DEBUG] [%s: %s: %d]%s\n", now.Format("2006-01-02 15:04:05"), funcName, fileName, lineNo, msg)
	}
}
func (l Logger) Fatal(msg string) {
	if l.Level <= FATAL {
		now := time.Now()
		funcName, fileName, lineNo := getInfo(2)
		fmt.Printf("[%s][DEBUG] [%s: %s: %d]%s\n", now.Format("2006-01-02 15:04:05"), funcName, fileName, lineNo, msg)
	}
}
