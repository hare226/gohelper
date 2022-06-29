package catLog

import (
	"fmt"
	"github.com/shiena/ansicolor"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var logMsgPool *sync.Pool // 异步时日志缓冲
const defaultLogChanLen = 1e3

var LOG *Logger

var levelMap = map[LogLevel]string{
	LevelDebug:   "Debug ",
	LevelInfo:    "Info ",
	LevelWarning: "Warning ",
	LevelError:   "Error ",
}

type Logger struct {
	lock       *sync.Mutex // modify  logger need lock
	level      LogLevel
	writer     *MultiWriter
	msgChanLen int64
	msgChan    chan *logMsg
}

func (object *Logger) write(m *logMsg) {
	headTime, _, _ := formatTimeHeader(m.when)
	color := formatColor(m.msg, m.level)
	object.writer.writeConsole(append(headTime, color...))
	mb := levelMap[m.level] + m.msg
	object.writer.writeFile(append(headTime, mb...))

}

// 清理一次日志信息
// @param crtTime 当前时间
func (object *Logger) clearOnce(crtTime time.Time) {
	dir, err := ioutil.ReadDir(catLogDir)
	if err != nil {
		_, _ = object.writer.console.Write([]byte(fmt.Sprintf("清理日志文件失败 目录: %s 错误: %s", catLogDir, err.Error())))
		return
	}
	crtSeconds := crtTime.Unix()
	for _, info := range dir {
		if (crtSeconds-info.ModTime().Unix()) > (86400*7) && info.IsDir() == false {
			delFile := filepath.Join(catLogDir, info.Name())
			if err := os.Remove(delFile); err != nil {
				_, _ = object.writer.console.Write([]byte(fmt.Sprintf("清理日志文件失败 文件: %s 错误: %s \n", delFile, err.Error())))
				continue
			}
			_, _ = object.writer.console.Write([]byte(fmt.Sprintf("删除日志文件  %s  \n", delFile)))
		}
	}
}

// 清理日志,目前默认按照时间来清理
func (object *Logger) clear() {
	timer := time.NewTimer(time.Second)
	for crtTime := range timer.C {
		timer.Reset(time.Hour * 10)
		object.clearOnce(crtTime)
	}
}

func (object *Logger) start() {
	for {
		select {

		// 切换日志
		case logName := <-ToggleLogic.notify:
			object.writer.fileWriter.Sync()
			object.writer.fileWriter.Close()
			object.writer.fileWriter = newLogFile(logName)

		// 写入日志
		case m := <-object.msgChan:
			object.write(m)
			logMsgPool.Put(m)
		}
	}
}

func NewWriter() *MultiWriter {
	return &MultiWriter{
		console:    ansicolor.NewAnsiColorWriter(os.Stdout),
		fileWriter: newLogFile(ToggleLogic.formatLogName("")),
	}
}

// 传入需要写文件的目录
func newCatLog() *Logger {

	return &Logger{
		lock:       new(sync.Mutex),
		level:      LevelDebug,
		writer:     NewWriter(),
		msgChanLen: defaultLogChanLen,
		msgChan:    make(chan *logMsg, defaultLogChanLen),
	}

}

// func formatLog(f interface{}, v ...interface{}) string {
//	var msg string
//	switch f.(type) {
//	case string:
//		msg = f.(string)
//		if len(v) == 0 {
//			return msg
//		}
//		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
//			//format string
//		} else {
//			//do not contain format char
//			msg += strings.Repeat(" %v", len(v))
//		}
//	default:
//		msg = fmt.Sprint(f)
//		if len(v) == 0 {
//			return msg
//		}
//		msg += strings.Repeat(" %v", len(v))
//	}
//	return fmt.Sprintf(msg, v...)
// }
func formatColor(s string, l LogLevel) string {
	var color string
	switch l {
	// case LevelDebug: //green
	//	color = fmt.Sprintf("%s", levelMap[l])
	case LevelInfo: // blue
		color = fmt.Sprintf("%s", levelMap[l])
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
	s = "[" + filename + ":" + strconv.Itoa(line) + "] " + s + "\n" //
	lm := logMsgPool.Get().(*logMsg)
	lm.level = l
	lm.msg = s
	lm.when = time.Now()
	object.msgChan <- lm
}

func Info(message string) {
	LOG.writeMsg(message, LevelInfo)
}
func Warning(message string) {
	LOG.writeMsg(message, LevelWarning)
}
func Error(message string) {
	LOG.writeMsg(message, LevelError)
}

// 初始化目录
func init() {
	LOG = newCatLog() // 初始化日志文件
	go LOG.start()
	go LOG.clear()
	// go LOG.autoToggleName()
	logMsgPool = &sync.Pool{ // 初始化日志pool
		New: func() interface{} {
			return &logMsg{}
		},
	}
}
