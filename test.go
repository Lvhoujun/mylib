package mylib

/*
* 测试模块
*
 */

import (

    //"fmt"

    "testing"

)

// ====================================================================
// Struct functions
// ====================================================================




// ====================================================================
// Api functions
// ====================================================================



// ====================================================================
// Internal functions
// ====================================================================



func TestEntry() {

    // fmt.Println("sync", testing.Benchmark(BenchmarkChannelSync).String())

    // fmt.Println("buffered",testing.Benchmark(BenchmarkChannelBuffered).String())

}

func BenchmarkChannelSync(b *testing.B) {

    ch := make(chan int)

    go func() {

        for i := 0; i < b.N; i++ {

            ch <- i

        }

        close(ch)

    }()

    for _ = range ch {

    }

}
