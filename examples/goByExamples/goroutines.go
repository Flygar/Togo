/*
Go 协程(goroutine) 在执行上来说是轻量级的线程。
goroutine是Go并行设计的核心。goroutine说到底其实就是线程，但是它比线程更小，十几个goroutine可能体现在底层就是五六个线程，Go语言内部帮你实现了这些goroutine之间的内存共享。执行goroutine只需极少的栈内存(大概是4~5KB)，当然会根据相应的数据伸缩。也正因为如此，可同时运行成千上万个并发任务。goroutine比thread更易用、更高效、更轻便。
goroutine是通过Go的runtime管理的一个线程管理器。goroutine通过go关键字实现了，其实就是一个普通的函数。

runtime.Gosched()//表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine。

同步：当前函数执行完了才继续往下执行
异步：统一起跑线同时执行，不必等待当前函数是否执行完
*/
package main

import (
	"fmt"
)

func fg(from string) {
	for i := 0; i < 10000000; i++ {
		fmt.Println(from, ":", i)
		//runtime.Gosched()
	}
}

func main() {
	go fg("hello")
	//time.Sleep(3000000000)

	// 假设我们有一个函数叫做 `fg(s)`。一般会这样同步(synchronously)调用；// 当前 Goroutines 执行
	fg("direct")

	// 使用 `go fg(s)` 在一个 Go 协程中调用这个函数。
	// 这个新的 Go 协程将会并发(concurrently)执行这个函数。
	// 开一个新的 Goroutines 执行
	go fg("goroutine")

	// 你也可以为匿名函数启动一个 Go 协程。
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 现在这两个 Go 协程在独立的 Go 协程中异步(asynchronously)运行，所以程序直接运行到这一行。这里的 `Scanln` 代码需要我们在程序退出前按下任意键结束。
	fmt.Scanln()
	fmt.Println("done")
	
}
