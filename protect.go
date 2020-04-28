// handler panic let programe keep run
package mylib

import (
	//"fmt"
	"runtime"
	"runtime/debug"
)


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
				ERROR("runtime error,panic:%v,stack:%s",err,string(debug.Stack()))
			default: // 非运行时错误
				//fmt.Println("error:", err)
		}
	}()
	entry()
}
