/*
通道(Channels) 是连接多个 Go 协程的管道。你可以从一个 Go 协程 将值发送到通道，然后在别的 Go 协程中接收。
默认情况下，channel 接收和发送数据都是阻塞的,不需要显式的 lock，是线程安全的
默认通道是 无缓冲 的，这意味着只有在对应的接收（<- chan） 通道准备好接收时，才允许进行发送（chan <-）。可缓存通道 允许在没有对应接收方的情况下，缓存限定数量的值
通道的方向：当使用通道作为函数的参数时，你可以指定这个通道是不是 只用来发送或者接收值。这个特性提升了程序的类型安全性
通道选择器：Go 的通道选择器 让你可以同时等待多个通道操作。 Go 协程和通道以及选择器的结合是 Go 的一个强大特性。
channel的超时处理：超时 对于一个连接外部资源，或者其它一些需要花费执行时间 的操作的程序而言是很重要的。得益于通道和 select，在 Go 中实现超时操作是简洁而优雅的。
非阻塞通道：常规的通过通道发送和接收数据是阻塞的。然而，我们可以 使用带一个 default 子句的 select 来实现非阻塞 的 发送、接收，甚至是非阻塞的多路 select。
channel的关闭：关闭 一个通道意味着不能再向这个通道发送值了。这个特性可以 用来给这个通道的接收方传达工作已经完成的信息。
channel的遍历： for range
在一个nil的channel上发送和接收操作会被永久阻塞。fatal error: all goroutines are asleep - deadlock!
channel 是 引用类型。 必须初始化才能写入数据，即make后才能使用。 channel是有类型的。
channel的cap不能自动增长
如果channel中的数据取完了再取的话会报错
channel 如果close了，那只能读，不能写。
gotoutine中使用recover可以解决某一协程panic的问题，可以有效避免直接从主线程退出
*/
package main

import (
	"fmt"
	"time"
)

// 声明channel
var intChan chan int
var mapChan chan map[int]string
var personChan chan Personor
var personChan chan *Person
// 推荐使用make来声明channel，跟简单易懂
fooChan := make(chan interface{},5) // 可以存放5哥任意类型的空接口的channel，



func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

// 这是一个我们将要在 Go 协程中运行的函数。`done` 通道将被用于通知其他 Go 协程这个函数已经工作完毕。
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	// 发送一个值来通知我们已经完工啦。
	done <- true
}

// channel的方向：`ping` 函数定义了一个只允许发送数据的通道。尝试使用这个通道来接收数据将会得到一个编译时错误。
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// channel的方向：`pong` 函数允许通道（`pings`）来接收数据，另一通道（`pongs`）来发送数据。
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	// 使用 `make(chan val-type)` 创建一个新的通道。
	// 通道类型就是他们需要传递值的类型。
	messages := make(chan string)

	// 使用 `channel <-` 语法 _发送(send)_ 一个新的值到通道中。这里我们在一个新的 Go 协程中发送 `"ping"` 到上面创建的`messages` 通道中。
	go func() { messages <- "ping" }()

	// 使用 `<-channel` 语法从通道中 _接收(receives)_ 一个值。这里将接收我们在上面发送的 `"ping"` 消息并打印出来。
	msg := <-messages
	fmt.Println(msg)

	//--------------------------------------------
	// channel的缓冲
	// 这里我们 `make` 了一个通道，最多允许缓存 2 个值。
	messages2 := make(chan string, 2)

	// 因为这个通道是有缓冲区的，即使没有一个对应的并发接收
	// 方，我们仍然可以发送这些值。
	messages2 <- "buffered"
	messages2 <- "channel"

	// 然后我们可以像前面一样接收这两个值。
	fmt.Println(<-messages2)
	fmt.Println(<-messages2)

	//--------------------------------------------
	//channel的同步
	// 运行一个 worker Go协程，并给予用于通知的通道。
	done := make(chan bool, 1)
	go worker(done)
	// 程序将在接收到通道中 worker 发出的通知前一直阻塞。如果你把 <- done 这行代码从程序中移除，程序甚至会在 worker 还没开始运行时就结束了。
	<-done // 可以把这个 动作（<-done） 赋值给一个变量，得到结果，一举两得。 _ <-done

	//--------------------------------------------
	//channel的方向
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

	//--------------------------------------------
	//channel的选择
	c1 := make(chan string)
	c2 := make(chan string)

	// 各个通道将在若干时间后接收一个值，这个用来模拟例如并行的 Go 协程中阻塞的 RPC 操作
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	// 我们使用 `select` 关键字来同时等待这两个值，并打印各自接收到的值。
	// 我们首先接收到值 "one"，然后就是预料中的 "two" 了。
	// 注意从第一次和第二次 Sleeps 并发执行，总共仅运行了 两秒左右。
	// 如果2个channel同时有值，select会随机选择一个执行
	// 没有default，这是阻塞接受。（一直阻塞到符合条件的case可以接受）
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	//--------------------------------------------
	//channel的超时处理
	// 在我们的例子中，假如我们执行一个外部调用，并在 2 秒后
	// 通过通道 `cg` 返回它的执行结果。
	cg := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		cg <- "result 1"
	}()

	// 这里是使用 `select` 实现一个超时操作。
	// `res := <- cg` 等待结果，`<-Time.After` 等待超时时间 1 秒后发送的值。由于 `select` 默认处理第一个已准备好的接收操作，如果这个操作超过了允许的 1 秒的话，将会执行超时 case。
	select {
	case res := <-cg:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	// 如果我允许一个长一点的超时时间 3 秒，将会成功的从 `cg2`接收到值，并且打印出结果。
	cg2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		cg2 <- "result 2"
	}()
	select {
	case res := <-cg2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}

	//--------------------------------------------
	//非阻塞通道
	messages3 := make(chan string)
	signals := make(chan bool)

	// 这里是一个非阻塞接收的例子。如果在 `messages3` 中存在，然后 `select` 将这个值带入 `<-messages3` `case`中。如果不是，就直接到 `default` 分支中。
	select {
	case msg3 := <-messages3:
		fmt.Println("received message3", msg3)
	default:
		fmt.Println("no message3 received")
	}

	// 一个非阻塞发送的实现方法和上面一样。
	msg3 := "hi"
	select {
	case messages3 <- msg3:
		fmt.Println("sent message3", msg3)
	default:
		fmt.Println("no message3 sent")
	}

	// 我们可以在 `default` 前使用多个 `case` 子句来实现一个多路的非阻塞的选择器。
	// 这里我们试图在 `messages`和 `signals` 上同时使用非阻塞的接收操作。
	select {
	case msg3 := <-messages3:
		fmt.Println("received message", msg3)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

	//--------------------------------------------
	//channel的关闭
	jobs := make(chan int, 5)
	done2 := make(chan bool)

	// 这是工作 Go 协程。使用 `j, more := <- jobs` 循环的从`jobs` 接收数据。在接收的这个特殊的二值形式的值中，
	// 如果 `jobs` 已经关闭了，并且通道中所有的值都已经接收完毕，那么 `more` 的值将是 `false`。当我们完成所有的任务时，将使用这个特性通过 `done2` 通道去进行通知。
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done2 <- true
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
	<-done2

	//--------------------------------------------
	//通道的遍历
	// 我们将遍历在 `queue` 通道中的两个值。
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// 这个 `range` 迭代从 `queue` 中得到的每个值。因为我们
	// 在前面 `close` 了这个通道，这个迭代会在接收完 2 个值
	// 之后结束。如果我们没有 `close` 它，我们将在这个循环中
	// 继续阻塞执行，等待接收第三个值
	for elem := range queue {
		fmt.Println(elem)
	}
	// 如果不close，如何实现： channel中有值就取，没值就等待有值？见下方实例2

	//----------------------------------------------
	// 使用recover可以解决某一协程panic的问题，可以有效避免直接从主线程退出
	func testGoroutineRecover() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recover,no worry")
			}
		}()
		for i := 0; i < 10; i++ {
			if i == 5 {
				panic("get 5")
			}
		}
	
	}
	
	func TestMainTest(t *testing.T) {
		go testGoroutineRecover()
		go testGoroutineRecover()
		go testGoroutineRecover()
	}

	//--------------------------------------------
	//实例1
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	//--------------------------------------------
	//实例2
	// 操作同一个管道，一边写，一边读，主线程需要等待2个goroutine协程都结束
	intChan :=make(chan int,50)
	isDone :=make(chan bool,1)

	go func() {
		for i:=0;i<cap(intChan);i++{
			intChan<-i+100
		}
		close(intChan)
	}()

	go func() {
		fmt.Println(len(intChan))
		for vChan :=range intChan{
			t.Log(vChan)
		}
		isDone<-true
	}()
	<-isDone

	//--------------------------------------------
	//实例3: 课堂练习（利用协程快速计算1-2000的和）
	// 协程1: 放入从1到2000的数字到channal中 （为毛不直接放到列表？因为下一步我要从协程中读取）
	// 额外创建4个协程从协程1中读取n个数字的数值并计算和后放到结束channel（或列表中），直到读光
	// 计算上一步的channel（或列表）中的所有值的和。



}
