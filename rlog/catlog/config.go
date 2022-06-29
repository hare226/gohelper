package catLog

import (
	"os"
	"path/filepath"
)

// 日志文件路径
var catLogDir string

func init() {
	rootPath, _ := filepath.Abs(filepath.Join(filepath.Dir(os.Args[0])))
	catLogDir = filepath.Join(rootPath, "log")
}

// 获取插件日志目录
func CatLogDir() string {
	return catLogDir
}

const (
	// PrintConsole bool = true
	PrintConsole bool = false
)
