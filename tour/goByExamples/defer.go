/*

Defer 被用来确保一个函数调用在程序执行结束前执行。同 样用来执行一些清理工作。 defer 用在像其他语言中的 ensure 和 finally用到的地方
关键字 defer 允许我们推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个语句或函数（为什么要在返回之后才执行这些语句？因为 return 语句同样可以包含一些操作，而不是单纯地返回某个值）。
defer 语句会将函数推迟到外层函数返回(return)之后执行。
推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
非常重要：defer是取当前环境的变量
*/
package main

import (
	"fmt"
	"os"
)

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}

//当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）
func fdefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d\n", i)
	}
}

//变量 ret 的值为 2，因为 ret++ 是在执行 return 1 语句后发生的
func fdefer2() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

func main() {
	f:=createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)

	//----------------------------------
	//后进先出
	fdefer()

	//----------------------------------
	fmt.Println("ret: ",fdefer2())
	
}
