package mylib

/*
* 协程惰性生成器
* 
*/

import (
	// "runtime"
	// "time"
)

// ====================================================================
// Struct functions
// ====================================================================

type Any interface{}
type EvalFunc func(Any)(Any,Any)
type FibonacciFunc func(uint64,uint64)(uint64,uint64)

// ====================================================================
// Api functions
// ====================================================================

func RoutineLazyEntry(){
	//GeneratorTest()
	FibonacciTest()
}

// ====================================================================
// Internal functions
// ====================================================================

func GeneratorTest(){
	eventFun := func(state Any) (Any,Any){
		os := state.(int)
		ns := os + 2
		return os,ns
	}
	even := EvalGenerator(eventFun, 0)
	for i := 0; i < 15; i++ {
		DEBUG("%dth even: %d",i,even())
	}
}

// @doc 高阶函数
func EvalGenerator(eventFun EvalFunc, state Any) func() Any{
	retChan := make(chan Any)
	loopFun := func(){
		var actState Any = state
		var retVal Any
		for{
			retVal,actState = eventFun(actState)
			retChan <- retVal
		}
	}
	go loopFun()
	retFun := func() Any {
		return <- retChan 
	}
	return retFun
}

// 斐波拉契懒惰计数器版
func FibonacciTest(){
	fiboFun := FibonacciFun()
	fiboGen := FibonacciGenerator(fiboFun)
	for i := 1; i < 50; i++ {
		DEBUG("%dth even: %d",i,fiboGen())		
	}
}

// @doc 斐波拉契生成函数
func FibonacciFun() func(x,y uint64) (uint64,uint64){
	Fun := func(x,y uint64) (uint64,uint64) {
		x2,y2:=y,x+y
		return x2,y2
	}
	return Fun
}

// @doc 斐波拉契生成器
func FibonacciGenerator(fibonacciFun FibonacciFunc) func() uint64{
	retChan := make(chan uint64)	
	loopFun := func(){
		var x uint64  = 0	
		var y uint64  = 1
		for{
			x,y = fibonacciFun(x,y)
			retChan <- x
		}
	}
	go loopFun()
	return func() uint64 {
		return <- retChan
	}
}