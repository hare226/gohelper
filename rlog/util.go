package rlog

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func formatDay() string {
	year, month, day := time.Now().Date()
	return fmt.Sprintf(DateFormat,
		strconv.FormatInt(int64(year), 10)[2:],
		int(month),
		day,
	)
}

func getBaseName() string {
	base := filepath.Base(os.Args[0])
	suffix := filepath.Ext(base)
	return strings.TrimSuffix(base, suffix)
}

func genLogName() string {
	return fmt.Sprintf(LogNameFormat, formatDay(), getBaseName())
}

func genLogFile() *os.File {
	rootPath, err := os.Getwd()
	if nil != err {
		panic("获取工作目录失败 " + err.Error())
	}

	logFilePath, err := filepath.Abs(filepath.Join(rootPath, LogRelativePath, genLogName()))
	if nil != err {
		panic("拼接日志文件目录失败 " + err.Error())
	}

	if err := os.MkdirAll(filepath.Dir(logFilePath), os.ModeDir); nil != err {
		panic("创建日志目录失败 " + err.Error())
	}

	fd, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if nil != err {
		panic("创建日志文件失败 " + err.Error())
	}

	return fd
}

func getNowTime() string {
	now := time.Now()
	h := now.Hour()
	m := now.Minute()
	s := now.Second()
	ms := now.UnixMilli() % 1000

	time := fmt.Sprintf(TimeFormat, h, m, s, ms)

	return fmt.Sprintf("[%s]", time)
}

func keepOneFolder(path string) string {
	folders := strings.Split(path, Separator)
	index := len(folders)
	if index < 2 {
		panic("程序工作路径过短")
	}
	return filepath.Join(folders[index-2], folders[index-1])
}

func getFileLine() (file string, line int) {
	_, file, line, _ = runtime.Caller(CallerDepth)
	file = keepOneFolder(file)

	// 通过pc获取被调用的函数信息
	// funcName = runtime.FuncForPC(pc).Name()
	return
}
