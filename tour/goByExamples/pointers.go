package main

import (
	"fmt"
)

/*
指针
带指针(或值)参数的函数必须接受一个指针(或值);
而以指针(或值)为接收者的方法被调用时，接收者既能为值又能为指针

值类型的变量的值存储在栈中
引用传递：slices，maps，channel；被引用的变量会存储在堆中，以便进行垃圾回收，且比栈拥有更大的内存空间。
字符串：值类型，且值不可改变
range:值拷贝
数组：值拷贝
切片：本质就是指向数组的指针，引用传递；所以不需要指针指向切片

为保证内存安全，golang中不允许做指针运算（与c的区别）
对一个空指针的反向引用是不合法的，并且会使程序崩溃

高级应用：你可以传递一个变量的引用（如函数的参数），这样不会传递变量的拷贝。
*/

func main() {
	//定义值类型变量i1，变量i1存储在栈中
	var i1 = 5
	//定义指针变量intP，指针变量intP存储在堆中
	var intP *int
	//指针变量intP指向i1的内存地址(存储5这个值的内存块的内存地址)；通过 &i 语法来取得 i 的内存地址，即指向 i 的指针。
	intP = &i1
	//这里的 * 号是一个类型更改器。将得到这个指针指向地址上所存储的值；这被称为反引用（或者内容或者间接引用）操作符；另一种说法是指针转移。使用一个指针引用一个值被称为间接引用。或者是解引用
	fmt.Println(*intP)
	//改变：对一个解引用的指 针赋值将会改变这个指针引用的真实地址的值。
	*intP = 10
	//变量i1所指向的内存中的值变成了10
	fmt.Println(i1)

// ------------------------------ 	

	num := 10
	fmt.Println("&num:", &num)
	whatPtr(&num)
	fmt.Println("num:", num)
}


func whatPtr(ptr *int) { //  接受一个指针（ `某个变量的地址` ）
	fmt.Println("ptr:", ptr)     //  返回 栈中 `新开辟内存` 上存放的值（存的值：上面接收的 `指针` , `某个变量的地址` ）
	fmt.Println("&ptr:", &ptr)   //  返回 栈中 `新开辟内存` 它自已的地址。 ps：它存放的值是 `指针` , `某个变量的地址`
	fmt.Println("*&ptr:", *&ptr) // 反引用，返回 `栈中新开辟内存它自已的地址` 上面的值，其实就是 `指针` , `某个变量的地址`
	fmt.Println("*ptr:", *ptr)   // 返回 （`指针` , `某个变量的地址`）上的值

	*ptr = (*ptr) + 10           // 对（指针` ,`某个变量的地址`） 上的值 进行+10 操作
	fmt.Println("ptr:", ptr)     // 还是 `某个变量的地址` （不变）
	fmt.Println("*ptr的值:", *ptr) // `某个变量的地址` 上的值（跟新了+10操作，返回20）

}


