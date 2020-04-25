// Package mylib implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package mylib

import (
	"fmt"
	"container/list"
)

// ====================================================================
// API functions
// ====================================================================

func EmptyFunc(){

}

func MapFuncExample(){
	MapFunc := map[int] func() int {
		1:func() int {return 10},
		2:func() int {return 20},
		3:func() int {return 30},
	}
	fmt.Println(MapFunc)
}


//正序打印
func print_list(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, "  ")
	}
	fmt.Println()
}

//倒序打印
func rprint_list(l *list.List) {
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, "  ")
	}
	fmt.Println()
}