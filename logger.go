// logger file handler
package mylib

import (
	"log"
	"os"

	//"fmt"
	"runtime"
	"path"	
	"strconv"
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

const DEFAULT_LL int = LL_DEBUG   //默认日志级别DEBUG

const STR_ERROR string = "Error"
const STR_WARNING string = "Warning"
const STR_DEBUG string = "Debug"

var LogLevel *int = nil
var errFile *os.File = nil


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
		Debug = log.New(errFile, "[D]", log.Ldate|log.Lmicroseconds)
	}
	if Warning == nil {
		Warning = log.New(errFile, "[W]", log.Ldate|log.Lmicroseconds)
	}
	if Error == nil {
		//Error = log.New(io.MultiWriter(os.Stderr, errFile), "[E]", log.Ldate|log.Lmicroseconds|log.Lshortfile)  //多端显示
		Error = log.New(errFile, "[E]", log.Ldate|log.Lmicroseconds) //多端显示
	}
}

//error级别写日志
func ERROR(format1 string, args ...interface{}) {
	format := CallInfo()+" "+format1
	if *LogLevel >= LL_ERROR {
		Error.Printf(format, args...)
	}
}

//debug级别写日志
func WARNING(format string, args ...interface{}) {
	if *LogLevel >= LL_WARNING {
		Warning.Printf(format, args...)
	}
}

//debug级别写日志
func DEBUG(format string, args ...interface{}) {
	if *LogLevel >= LL_DEBUG {
		Debug.Printf(format, args...)
	}
}

//调用文件与行号,eg main.go:18
func CallInfo() string {
	_, var2, var3, _ := runtime.Caller(2)
	nameWithSuffix:=path.Base(var2)
	return nameWithSuffix+":"+strconv.Itoa(var3)
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
