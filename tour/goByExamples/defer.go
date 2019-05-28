/*

1. Defer 被用来确保一个函数调用在程序执行结束前执行。同样用来执行一些清理工作。 defer 用在像其他语言中的 ensure 和 finally用到的地方
2. 关键字 defer 允许我们推迟到函数返回之前（return 语句之后）一刻才执行某个语句或函数（为什么要在返回之后才执行这些语句？因为 return 语句同样可以包含一些操作，而不是单纯地返回某个值）。
3. defer 语句会将函数推迟到外层函数返回(return)之后执行。
4. 推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
5. 非常重要：defer是取当前环境的变量
6. defer将语句放入到栈时（值拷贝），也会将相关的值拷贝同时入栈
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

//3. 重要：变量 ret 的值为 3，因为 ret++ 是在执行 return 2 语句后发生的.(ret 的值不是2 )
func fdefer2() (ret int) {
	defer func() {
		ret++
	}()
	return 2
}

//4. 当有多个 defer 行为被注册时，它们会以逆序执行（即后进先出）
func fdefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d\n", i)
	}
}

//6. defer将语句放入到栈时，也会将相关的值拷贝同时入栈（值拷贝）
func sum(n1, n2 int) int {
	defer fmt.Println("n1's value: ", n1)
	defer fmt.Println("n2's value:", n2)

	n1++
	n2++

	ret := n1 + n2
	fmt.Println("ret's value:", ret)

	return ret

}

func main() {
	f:=createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)

	//----------------------------------

	//3. 返回3
	fmt.Println("ret: ",fdefer2())

	//6. 值拷贝
	/*
	ret's value: 32
	n2's value: 20
	n1's value: 10
	32
	*/
	ret := sum(10, 20)
	fmt.Println(ret)
	
}
