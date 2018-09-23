package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// 这是工作 Go 协程。使用 `j, more := <- jobs` 循环的从`jobs` 接收数据。在接收的这个特殊的二值形式的值中；
	// 如果 `jobs` 已经关闭了，并且通道中所有的值都已经接收完毕，那么 `more` 的值将是 `false`。当我们完成所有的任务时，将使用这个特性通过 `done` 通道去进行通知。
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// 这里使用 `jobs` 发送 3 个任务到工作函数中，然后关闭 `jobs`。
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// 我们使用前面学到的[通道同步](../channel-synchronization/)方法等待任务结束。
	<-done
}
