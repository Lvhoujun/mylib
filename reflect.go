package mylib

//package main

import (
	"fmt"
	"reflect"
)

// ====================================================================
// struct defination
// ====================================================================

// ====================================================================
// API defination
// ====================================================================

func TestReflect() {
	var x float64 = 3.4
	fmt.Printf("TypeOf x=%v\n", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	// fmt.Printf("v=%v\n", v)
	// fmt.Printf("v.Type=%v\n", v.Type())
	// fmt.Printf("v.Kind=%v\n", v.Kind())	
	// fmt.Println("value:", v.Float())
	// fmt.Println("value Int:", int(v.Float()))  //v的类型是float64，这里不能用v.Int,会触发 reflect: call of reflect.Value.Elem on float64 Value 错误
	// fmt.Println("Interface:", v.Interface().(float64))
	v = v.Elem()
	fmt.Printf("CanSet=%v\n",v.CanSet())
	fmt.Printf("v=%v\n",v)	
	fmt.Printf("x=%v\n",x)	
}
