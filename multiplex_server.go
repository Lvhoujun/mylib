package mylib

/*
* 多路复用服务器模型
*
 */

import (
	"time"
	// "runtime"
	//"time"
)

// ====================================================================
// Struct functions
// ====================================================================

const MAXREQS = 10 //同时处理的最大并发数
const TASK_QUEUE_LEN = 50  //任务数

var sem chan int = make(chan int, MAXREQS)

type request struct {
	a, b   int
	replyc chan int //请求回复通道
}

type opFun func(a, b int) int

// ====================================================================
// Api functions
// ====================================================================

func MultiServerEntry() {
	MultiServerTest()
}

// ====================================================================
// Internal functions
// ====================================================================

/*
 * 任务队列通道 service 缓冲必须要大于等于最大任务数
 * sem <- 1 控制了同时并发的最大协程数，任务队列接收会阻塞，故接收池子要足够大
 * 此测试Run sleep 1s,sem缓冲满以达测试目的
 * 
 */

func MultiServerTest() {
	handleFun := func(a, b int) int {
		return a + b
	}
	service := StartServer(handleFun)	
	var reqs [TASK_QUEUE_LEN]request
	for i := 0; i < TASK_QUEUE_LEN; i++ {

		req := &reqs[i]
		req.a = i
		req.b = i + TASK_QUEUE_LEN
		req.replyc = make(chan int)
		service <- req
	}
	for i := 0; i < TASK_QUEUE_LEN; i++ {
		DEBUG("reqs[%d]:a=%d,b=%d,a+b=%d", i, reqs[i].a, reqs[i].b, <-reqs[i].replyc)
	}
}

func Run(fun opFun, req *request) {
	time.Sleep(time.Second)
	req.replyc <- fun(req.a, req.b)
	<-sem
}

// server协程:接收到请求后开启一个协程处理请求
func Server(fun opFun, service chan *request) {
	for {
	
		//接受请求
		//当sem缓冲满了，此处会阻塞，达到限制最大并发的目的
		sem <- 1

		req := <-service
		//协程处理请求
		go Run(fun, req)
	}
}

//开启server 开启一个协程用于接收请求
func StartServer(fun opFun) chan *request {
	service := make(chan *request, TASK_QUEUE_LEN)
	go Server(fun, service)
	return service
}
