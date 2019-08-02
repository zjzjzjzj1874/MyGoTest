// 简单的日志打印方法
package util

import (
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	lock sync.Mutex
	Config *LogConfig
)
const (
	InitVersion		= iota
	LogMaxSize		= 200	// 单个文件最大20M
	ErrLogPath		= "./log/errLog"	// TODO 读取配置文件
	InfoLogPath		= "./log/infoLog"	// TODO 读取配置文件
	Suffix			= ".log"			// TODO 读取配置文件
)
// 日志类型
const(
	LogTypeInfo = iota	// 普通日志
	LogTypeError		// 错误日志
)

type LogConfig struct{
	LogSize		int		// 日志大小
	LogSuffix	string	// 日志后缀
	LogPath		string	// 日志路径
	LogType		int		// 日志类型
	LogDate		string	// 日志日期
	LogVersion	int		// 日志版本
}

func init()  {
	Config = new(LogConfig)
	Config.LogSuffix = Suffix
	Config.LogDate = time.Now().Format("2006-01-02 15:04:05")
	Config.LogVersion = InitVersion
}

// 错误日志记录
func (l *LogConfig) Error(i... interface{}) {
	l.returnFileName("./log/errLog")
	fMode := os.O_APPEND|os.O_WRONLY
	lock.Lock()
	f, _ := os.OpenFile(Config.LogPath,fMode,0666)
	defer f.Close()
	_,_ = f.WriteString(time.Now().Format("2006-01-02 15:04:05")+ " [ERROR] | " + InterfaceToString(i) + "\n")
	lock.Unlock()
}

// 普通日志记录
func (l *LogConfig) Info(i... interface{}) {
	l.returnFileName("./log/infoLog")
	fMode := os.O_APPEND|os.O_WRONLY
	lock.Lock()
	f, _ := os.OpenFile(Config.LogPath,fMode,0666)
	defer f.Close()
	_,_ = f.WriteString(time.Now().Format("2006-01-02 15:04:05")+ " [INFO] | " + InterfaceToString(i) + "\n")
	lock.Unlock()
}

// 创建文件,判断大小,返回文件名
func (l *LogConfig) returnFileName(logPath string) {
	switch {
	case logPath == ErrLogPath:
		logPath = ErrLogPath
	default:
		logPath = InfoLogPath
	}
	Config.LogPath = logPath + Config.LogDate + "V" + strconv.Itoa(Config.LogVersion) + Config.LogSuffix
	if !l.isExist(){
		_ = l.create()
	}else{
		// 已经存在,就去判断文件大小
	}
}

// 创建文件
func (l *LogConfig) create() error {
	_, err := os.Create(Config.LogPath)
	if err != nil {
		return err
	}else{
		return nil
	}
}
// 文件（夹）是否存在
func (l *LogConfig) isExist() bool {
	_, err := os.Stat(Config.LogPath)
	return err == nil || os.IsExist(err)
}