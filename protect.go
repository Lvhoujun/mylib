// handler panic let programe keep run
package mylib

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

// ====================================================================
// struct defination
// ====================================================================

var StdErr bool = true //发生错误是否写标准输出

// ====================================================================
// API defination
// ====================================================================

// @doc 运行时错误是否写标准输出
// @param flag true要写
func SetStdErrWrteFlag(flag bool) {
	StdErr = flag
}

// 崩溃时需要传递的上下文信息
type PanicContext struct {
	Function string // 所在函数
}

func RunWithProtect(entry func()) {
	// 延迟处理的函数
	defer func() {
		// 发生宕机时，获取panic传递的上下文并打印
		err := recover()
		switch err.(type) {
		case runtime.Error: // 运行时错误
			//fmt.Printf("runtime err")
			stack := string(debug.Stack())
			ERROR("runtime error,panic:%v,stack:%s", err, stack)
			write_std_error(err,stack)
		default: // 非运行时错误			
			if err != nil {
				stack := string(debug.Stack())
				ERROR("runtime error,panic:%v,stack:%s", err, stack)
				write_std_error(err,stack)
			}
		}
	}()
	entry()
}

// ====================================================================
// Internal defination
// ====================================================================

func write_std_error(err interface{}, stack string){
	if StdErr==true{
		fmt.Println("runtime error,panic:%v,stack:%s", err, stack)
	}
}
