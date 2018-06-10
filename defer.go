package main

import (
	"fmt"
	"time"
)

/*
记录有关defer的坑
非常重要：defer是取当前环境的变量
在 main 函数中写一个用于打印 Hello World 字符串的匿名函数并赋值给变量 fv，然后调用该函数并打印变量 fv 的类型。
*/

func main() {
	fv := func() {
		fmt.Println("Hello World")
	}
	fv()
	fmt.Printf("Type of fv : %T", fv)

	func(u string) {
		fmt.Println(u)
	}("\nHello World2")

	//test1
	func() {
		fmt.Println("test1")
	}()

	//defer1
	fmt.Println("f() return:", f())

	//defer2
	f2()

	//defer22
	defer f22(time.Now())
	fmt.Println("1 ", time.Now())
	time.Sleep(5e9)

	//defer3
	fmt.Println(f3())

	//defer4
	fmt.Println(f4())

	//defer5
	fmt.Println(f5())

	//defer6
	f6()

	//defer7
	f7()

}

//defer1
func f() (ret int) {
	fmt.Println("----- defer1 -----")
	ret = 1
	defer func() {
		ret++
		fmt.Printf("defer1 ret=%d\n", ret)
	}()
	return 2
}

//defer2
func f2() {
	fmt.Println("----- defer2 func a() -----")
	i := 0
	i++
	t := time.Now()
	// defer声明时会先计算确定参数的值，defer推迟执行的仅是其函数体。
	defer fmt.Println("defer ", i, t)
	time.Sleep(5e9)
	i++
	fmt.Println("a() i=", i, time.Now())
	return
}

//defer2-2
func f22(t time.Time) {
	fmt.Println("----- defer22 -----")
	fmt.Println("2 ", t)
	fmt.Println("3 ", time.Now())
}

//defer3
func f3() (result int) {
	fmt.Println("----- defer3 -----")
	//改变了返回值的结果
	defer func() {
		result++
		fmt.Printf("defer %d", result)
	}()
	return 6
}

//defer4
func f4() (r int) {
	fmt.Println("----- defer4 -----")
	t := 5
	//改变的是t，并不是r
	defer func() {
		t = t + 5
	}()
	//5已经返回给r
	return t
}

//defer5
func f5() (r int) {
	fmt.Println("----- defer5 -----")
	//此r非彼r，可以把defer中的r改成其他形参
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func f6() {
	fmt.Println("----- defer6 -----")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func f7() {
	fmt.Println("----- defer7 -----")
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}
