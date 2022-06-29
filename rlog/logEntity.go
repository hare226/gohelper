package rlog

import (
	"fmt"
	"os"
)

// LogEntity 自定义日志结构体
type LogEntity struct {
	file *os.File
	msg  string
}

func (this *LogEntity) genMsg(level LogLevel, con string) {

	file, line := getFileLine()
	this.msg = fmt.Sprintf(MsgFromat, getNowTime(), level, file, line, con)
}

func (this *LogEntity) writeToConsole(con string) {
	if printToConsole {
		fmt.Println(con)
	}
}

func (this *LogEntity) writeToFile(con string) {
	if printToFile {
		this.file.Write([]byte(con))
	}
}

func (this *LogEntity) multiWrite() {
	this.writeToConsole(this.msg)
	this.writeToFile(this.msg)
}
