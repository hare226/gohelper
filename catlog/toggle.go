package catLog

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// 日志切换管理
type toggle struct {
	logName       string
	notify        chan string
	mutex         *sync.Mutex
	disableToggle bool
}

// 初始化
func (this *toggle) init() *toggle {
	this.notify = make(chan string, 1)
	this.mutex = new(sync.Mutex)
	go this.autoToggle()
	return this
}

// 获取一个日期的第一秒时间
// @param timestamp 时间戳
func (this *toggle) firstDaySecond(timestamp int64) time.Time {
	un := time.Unix(timestamp, 0).Format("2006-01-02") + " 00:00:00"
	location, e := time.ParseInLocation("2006-01-02 00:00:00", un, time.Local)
	if e != nil {
		return time.Time{}
	}
	return location
}

// 设置日志名称
// @param logName 日志名称
func (this *toggle) SetLogName(logName string) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if this.disableToggle = false; logName != "" {
		this.disableToggle = true
	}
	this.logName = this.formatLogName(logName)
	this.notify <- this.logName
}

// 获取日志名称
func (this *toggle) LogName() string {
	return this.logName
}

// 获取当前日志路径
func (this *toggle) LogPath() string {
	return filepath.Join(catLogDir, this.logName)
}

// 格式化日志名称
// @param logName 日志名称
func (this *toggle) formatLogName(logName string) string {
	if logName == "" {
		logName = time.Now().Format("06-01-02")
	}
	fileName := os.Args[0]
	fileName = strings.TrimSuffix(filepath.Base(fileName), filepath.Ext(fileName))
	return fmt.Sprintf("%s-%s.log", fileName, logName)
}

// 实现自动切换
func (this *toggle) autoToggle() {

	current := time.Now().Unix()
	ticker := time.NewTimer(time.Second * time.Duration((this.firstDaySecond(current).Unix()+86400)-current))

	for range ticker.C {
		if this.disableToggle == false {
			this.SetLogName("")
		}
		ticker.Reset(time.Second * 86400)
	}
}

// 日志全局对象
var ToggleLogic = new(toggle).init()
