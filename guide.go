package mylib

/*
* 常用错误编程习惯
*
*/

import (
	"fmt"
	// "runtime"
	"time"
)

// ====================================================================
// Struct functions
// ====================================================================



// ====================================================================
// Api functions
// ====================================================================

// func GuideEntry() {
// 	var remember bool = false
// 	if 1>0 {
// 		remember := true
// 		DEBUG("remember=%v",remember)
// 	}
// 	DEBUG("remember=%v",remember) //输出false,而不是true
// }

func GuideEntry(){
	values := [5]int{10,11,12,13,14}
	
	for ix := range values { // ix是索引值
        func() {
            fmt.Print(ix, "  ")
        }()        
    }
    fmt.Println()
    for ix := range values { // ix是索引值
        go func() {
            fmt.Print(ix, "  ")
        }() // 调用闭包打印每个索引值        
    }
    time.Sleep(time.Second)
    fmt.Println()
    for ix := range values { // ix是索引值
        go func(ix int) {
            fmt.Print(ix, "  ")
        }(ix) // 调用闭包打印每个索引值
        
    }

    time.Sleep(time.Second)
    fmt.Println()
    for ix := range values {
        val := values[ix]
        go func() {
            fmt.Print(val, " ")
        }()
    }
    
    
    time.Sleep(time.Second)
}

// ====================================================================
// Internal functions
// ====================================================================



var keyChar = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenKey(n int) string {
	if n == 0 {
		return string(keyChar[0])
	}
	l := len(keyChar)
	s := make([]byte, 20) // FIXME: will overflow. eventually.
	i := len(s)
	for n > 0 && i >= 0 {
		i--
		j := n % l
		n = (n - j) / l
		s[i] = keyChar[j]
	}
	return string(s[i:])
}