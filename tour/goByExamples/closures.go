package main

import "fmt"

/*
闭包：内层函数引用了外层函数中的变量，变量被引用后，它所在的函数结束，这变量也不会马上被烧毁。
golang中函数不能嵌套函数，但是可以在一个函数中包含匿名函数
*/

func main() {
	//c1 跟 c2 引用的是不同的环境，在调用 i++ 时修改的不是同一个 i，因此两次的输出都是 1
	c1 := f(0)
	c2 := f(0)
	c1()
	c1() // 引用c1相同的环境
	c2()

	fmt.Println("-----------")

	fmt.Println(intSeq()()) // 1
	fmt.Println(intSeq()()) // 1
	fmt.Println(intSeq()()) // 1

	test := intSeq()
	fmt.Println(test()) // 1
	fmt.Println(test()) // 2
	fmt.Println(test()) // 3

}

// 闭包函数0
func f(i int) func() int {
	return func() int {
		i++
		fmt.Println(i)
		return i
	}
}

// 闭包函数1
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// 闭包函数2
func outer(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
