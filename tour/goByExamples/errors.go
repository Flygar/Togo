/*
错误处理
符合 Go 语言习惯的做法是使用一个独立、明确的返回值来传递错误信息。
这与使用异常(exception)的 Java 和 Ruby以及在 C 语言中有时用到的重载(overloaded)的单返回/错误值有着明显的不同。
Go 语言的处理方式能清楚的知道哪个函数返回了错误，并能像调用那些没有出错的函数一样调用。
//defer panic recover

*/
package main

import (
	"errors"
	"fmt"
)

func main() {

	// 下面的两个循环测试了各个返回错误的函数。注意在 `if`行内的错误检查代码，在 Go 中是一个普遍的用法。
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// 你如果想在程序中使用一个自定义错误类型中的数据，你需要通过类型断言来得到这个错误类型的实例。
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

}

//使用defer + recover。 后续代码继续执行
func test() {
	defer func() {
		 //内置函数，可以捕获到异常
		if err:=recover();err!=nil{ //说明捕获到异常
			fmt.Println("err:",err)
			//发送邮件给 管理员
			fmt.Println("已发送邮件给XXXX")
		}
	}()
	num1:=10
	num2:=0
	ret:=num1/num2
	fmt.Println("ret:",ret)
}


//使用errors.New 和 panic(内置函数，接收空接口类型) 自定义错误(errors.New)+中断(panic)并退出程序
//函数去读取配置文件init.conf;如果文件名传入不正确，返回一个自定义的错误
func readConf(fileName string) (err error) {
	if fileName=="config.ini"{
		//读取
		return nil
	}else {
		//返回一个自定义的错误
		return  errors.New("自定义错误：读取文件失败")
	}
	//这里不需要return
}

func test02() {
	err:=readConf("config.ini")
	if err!=nil{
		//如果读取文件错误，就输出这个错误，并终止程序
		panic(err)
	}
}


func f1(arg int) (int, error) {
	if arg == 42 {

		// `errors.New` 构造一个使用给定的错误信息的基本`error` 值。
		return -1, errors.New("can't work with 42")

	}

	// 返回错误值为 nil 代表没有错误。
	return arg + 3, nil
}

// 通过实现 `Error` 方法来自定义 `error` 类型是可以的。
// 这里使用自定义错误类型来表示上面例子中的参数错误。
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// 在这个例子中，我们使用 `&argError` 语法来建立一个新的结构体，并提供了 `arg` 和 `prob` 这两个字段的值。
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}