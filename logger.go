// logger file handler
package mylib

import (
	"log"
	"os"

	"fmt"
	"io"
	// "path"
	// "runtime"
	// "strconv"
)

// ====================================================================
// API functions
// ====================================================================

var (
	Debug   *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

const (
	LL_ERROR   = iota + 1 //错误级别
	LL_WARNING            //警告级别
	LL_DEBUG              //debug级别
)

const DEFAULT_LL int = LL_DEBUG //默认日志级别DEBUG

const STR_ERROR string = "Error"
const STR_WARNING string = "Warning"
const STR_DEBUG string = "Debug"

var LogLevel *int = nil
var errFile *os.File = nil
var MultWrite bool = true //是否开启多端显示
var Depth = 2             //堆栈深度

func GetLogLevel() int {
	if LogLevel == nil {
		SetLogLevel(DEFAULT_LL)
	}
	return *LogLevel
}

func SetLogLevel(Level int) {
	var err error
	if errFile == nil {
		errFile, err = os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("打开日志文件失败：", err)
		}
	}
	LogLevel = new(int)
	*LogLevel = Level
	if Debug == nil {
		if MultWrite {
			Debug = log.New(io.MultiWriter(os.Stderr, errFile), "[D]", log.Ldate|log.Lmicroseconds|log.Lshortfile) //多端显示
		} else {
			Debug = log.New(io.MultiWriter(errFile), "[D]", log.Ldate|log.Lmicroseconds|log.Lshortfile) //多端显示
		}
	}
	if Warning == nil {
		if MultWrite {
			Warning = log.New(io.MultiWriter(os.Stderr, errFile), "[W]", log.Ldate|log.Lmicroseconds|log.Lshortfile) //多端显示
		} else {
			Warning = log.New(errFile, "[W]", log.Ldate|log.Lmicroseconds) //只写文件
		}
	}
	if Error == nil {
		if MultWrite {
			Error = log.New(io.MultiWriter(os.Stderr, errFile), "[E]", log.Ldate|log.Lmicroseconds|log.Lshortfile) //多端显示
		} else {
			Error = log.New(errFile, "[E]", log.Ldate|log.Lmicroseconds) //只写文件
		}
	}
}

//error级别写日志
func ERROR(format string, args ...interface{}) {
	if *LogLevel >= LL_ERROR {
		Error.Output(Depth, fmt.Sprintf(format, args...))
	}
}

//Warning级别写日志
func WARNING(format string, args ...interface{}) {
	if *LogLevel >= LL_WARNING {
		Warning.Output(Depth, fmt.Sprintf(format, args...))
	}
}

//debug级别写日志
func DEBUG(format string, args ...interface{}) {
	if *LogLevel >= LL_DEBUG {
		Debug.Output(Depth, fmt.Sprintf(format, args...))
	}
}

//根据日志级别获取日志描述串
func LogLevelString(LogLevel int) string {
	switch LogLevel {
	case LL_ERROR:
		return STR_ERROR
	case LL_WARNING:
		return STR_WARNING
	case LL_DEBUG:
		return STR_DEBUG
	default:
		return "unkown log_level"
	}
}

// ====================================================================
// inner functions
// ====================================================================
