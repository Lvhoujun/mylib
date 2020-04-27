// Package mylib implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package mylib

import (
	"fmt"
	"container/list"
	"errors"
)

// ====================================================================
// struct defination
// ====================================================================

//包体外可以直接访问
type Rectangle struct {
	Length int32         //包体外结构体可以访问
	Width int32          //包体外结构体可以访问
	private int32        //包体外结构体不可以访问
}

//包体外不能访问,只能通过工厂方法NewSquare初始化返回
type square struct {
	width int32
}


// ====================================================================
// API functions
// ====================================================================

func NewSquare(width int32) *square{
	s := new(square)
	s.width=width
	return s
}

func GetSqureWidth(s *square) (int32,error) {
	if s==nil {
		return 0,errors.New("nil pointer")
	}
	return (*s).width,nil
}

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