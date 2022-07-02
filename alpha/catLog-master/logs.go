package catLog

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/shiena/ansicolor"
)

var logMsgPool *sync.Pool // 异步时日志缓冲
const defaultLogChanLen = 1e3
const LogDir = "log"
const LogExtensionName = ".log"

var LOG *Logger

var levelMap map[LogLevel]string = map[LogLevel]string{
	LevelDebug:   "Debug ",
	LevelInfo:    "Info ",
	LevelWarning: "Warning ",
	LevelError:   "Error ",
}

func InitLog(dir string) {
	LOG = newCatLog(dir) // 初始化日志文件
	go LOG.start()
	logMsgPool = &sync.Pool{ // 初始化日志pool
		New: func() interface{} {
			return &logMsg{}
		},
	}
}

type Logger struct {
	lock       *sync.Mutex // modify  logger need lock
	level      LogLevel
	logDir     string
	writer     *MultiWriter
	msgChanLen int64
	msgChan    chan *logMsg
}

func (object *Logger) GetLogDir(m *logMsg) string {
	return object.logDir
}

func (object *Logger) write(m *logMsg) {
	headTime, _, _ := formatTimeHeader(m.when)
	color := formatColor(m.msg, m.level)
	object.writer.writeConsole(append(headTime, color...))
	// mb := "  " + levelMap[m.level] + "  " + m.msg
	mb := levelMap[m.level] + m.msg
	object.writer.writeFile(append(headTime, mb...))

}

// 清理日志,目前默认按照时间来清理
func (object *Logger) clear() {

}
func (object *Logger) start() {
	for {
		select {
		case m := <-object.msgChan: // 获取到了日志
			info, err := object.writer.fileWriter.Stat()
			if nil != err {
				object.writer.fileWriter.Sync()
				object.writer.fileWriter.Close()
				object.writer.fileWriter = GetNewLogFile(object.logDir)
			} else {
				if info.Size() >= DefaultFileSize {
					object.writer.fileWriter.Sync()
					object.writer.fileWriter.Close()
					object.writer.fileWriter = GetNewLogFile(object.logDir)
				}
			}
			object.write(m)
			logMsgPool.Put(m) // 回收日志
		}
	}
}

func NewWriter(dirPath string) *MultiWriter {
	return &MultiWriter{
		console:    ansicolor.NewAnsiColorWriter(os.Stdout),
		fileWriter: GetNewLogFile(dirPath),
	}
}

// 传入需要写文件的目录
func newCatLog(dirPath string) *Logger {
	if dirPath == "" {
		tmpDir, err := filepath.Abs(os.Args[0])
		if nil != err {
			panic(err)
		} else {
			dirName, _ := filepath.Split(tmpDir)
			dirPath = dirName
		}
	}
	return &Logger{
		lock:       new(sync.Mutex),
		level:      LevelDebug,
		logDir:     dirPath,
		writer:     NewWriter(dirPath),
		msgChanLen: defaultLogChanLen,
		msgChan:    make(chan *logMsg, defaultLogChanLen),
	}

}

func formatLog(f interface{}, v ...interface{}) string {
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
		if len(v) == 0 {
			return msg
		}
		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
			// format string
		} else {
			// do not contain format char
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}
	return fmt.Sprintf(msg, v...)
}
func formatColor(s string, l LogLevel) string {
	var color string
	switch l {
	case LevelDebug: // green
		color = fmt.Sprintf("\x1b[32m\x1b[1m%s\x1b[0m", levelMap[l])
	case LevelInfo: // blue
		color = fmt.Sprintf("\x1b[34m\x1b[1m%s\x1b[0m", levelMap[l])
	case LevelWarning:
		color = fmt.Sprintf("\x1b[33m\x1b[1m%s\x1b[0m", levelMap[l])
	case LevelError:
		color = fmt.Sprintf("\x1b[31m\x1b[1m%s\x1b[0m", levelMap[l])

	}
	return color + s
}

func (object *Logger) writeMsg(s string, l LogLevel) {
	if object.level > l { // 丢弃
		return
	}
	_, file, line, ok := runtime.Caller(2) // 获取文件调用的行
	if !ok {
		file = "???"
		line = 0
	}
	_, filename := path.Split(file)
	s = "[" + filename + ":" + strconv.Itoa(line) + "]" + s + "\n" //
	lm := logMsgPool.Get().(*logMsg)
	lm.level = l
	lm.msg = s
	lm.when = time.Now()
	object.msgChan <- lm
}
func Debug(f interface{}, v ...interface{}) {
	LOG.writeMsg(formatLog(f, v...), LevelDebug)
}
func Info(f interface{}, v ...interface{}) {
	LOG.writeMsg(formatLog(f, v...), LevelInfo)
}
func Warning(f interface{}, v ...interface{}) {
	LOG.writeMsg(formatLog(f, v...), LevelWarning)
}
func Error(f interface{}, v ...interface{}) {
	LOG.writeMsg(formatLog(f, v...), LevelError)
}

func GetLOGDir() string {
	return LOG.logDir
}

func Panic(f interface{}, v ...interface{}) {
	panic(formatLog(f, v...))
}
