package mylib

import (
	"math/rand"
	"runtime"
	"time"
)

// ====================================================================
// Struct functions
// ====================================================================

// ====================================================================
// Api functions
// ====================================================================

func RoutineTest() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//RoutineSelect()
	//fibonacci_select()
	//tick_channel_test()
	//timeout_test()
	StopChannelTest()
	DEBUG("main finish")
}

func RoutineSelect() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go producer1(ch1)
	go producer2(ch2)
	go consumer(ch1, ch2)
	time.Sleep(2 * time.Millisecond)
}

// ====================================================================
// Internal functions
// ====================================================================

// @do 多通道关闭
func StopChannelTest(){
	stopCh := make(chan int)
	workCh := make(chan int)
	Worker := func(i int, stopCh <-chan int){
		defer DEBUG("worker %d exit", i)
		DEBUG("worker %d start", i)
		for{
			select{
			case <-stopCh:
				return;
			case recv := <-workCh:
				DEBUG("worker %d recv %d",i,recv)	
			}
		}
	}

	for i := 1; i <= 2; i++ {
		go Worker(i,stopCh)
	}
	time.Sleep(2 * time.Second)
	workCh <- 10
	close(stopCh)
	time.Sleep(2 * time.Second)
}

//函数调用超时判断
func timeout_test() {
	ch := make(chan int, 1)
	timeout := 500 * time.Millisecond
	go func() {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		Random := r.Intn(1000)
		DEBUG("sleep %d ms",Random)
		Ms := time.Duration(Random) * time.Millisecond
		time.Sleep(Ms)
		ch <- Random
	}()
	select {
	case Random := <-ch:
		DEBUG("exeute succ Random=%d", Random)
	case <-time.After(timeout*time.Millisecond):
		DEBUG("time out=%d ms", timeout)
	}
}

// tick with channel
func tick_channel_test() {
	tick := time.Tick(1000 * time.Millisecond)
	boom := time.After(5 * time.Second)
	for {
		select {
		case <-tick:
			DEBUG("tick.")
		case <-boom:
			DEBUG("boom")
			return
		}
	}
}

func fibonacci_select() {
	ch := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			DEBUG("%d", <-ch)
		}
		quit <- 0
	}()
	fibonacci_routine(ch, quit)
}

// @doc协程版fibonacci
func fibonacci_routine(ch, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			return
		}
	}
}

// @doc 消费者
func consumer(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			DEBUG("ch1 recv %d", v)
		case v := <-ch2:
			DEBUG("ch2 recv %d", v)
		}
	}
}

// @doc 生产者1
func producer1(ch chan<- int) {
	for i := 0; ; i = i + 2 {
		ch <- i
	}
}

// @doc 生产者1
func producer2(ch chan<- int) {
	for i := 1; ; i = i + 2 {
		ch <- i + 5
	}
}

//@doc 只发
func OnlySendChan(ch chan<- int) {
	for {
		ch <- 1
	}
}

//@doc 只收
func OnlyRecvChan(ch <-chan int) {
	for value := range ch {
		DEBUG("value=%d", value)
	}
}

func Sum(a, b int, ch chan int) {
	ch <- a + b
}

func ChanSendData(ch chan string) {
	ch <- "this"
	ch <- "is"
	ch <- "send"
	ch <- "data"
	close(ch)
}

func ChanGetData(ch chan string) {
	for value := range ch {
		DEBUG("value=%s", value)
	}
}

func longwait() {
	DEBUG("longwait begin...")
	time.Sleep(5 * time.Second)
	DEBUG("longwait end...")
}

func shortwait() {
	DEBUG("shortwait begin...")
	time.Sleep(2 * time.Second)
	DEBUG("shortwait end...")
}
