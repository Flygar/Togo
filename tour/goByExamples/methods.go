/*
方法
>1. 带指针(或值)参数的函数必须接受一个指针(或值);
>2. 而以指针(或值)为接收者的方法被调用时，接收者既能为值又能为指针
>3. 任何类型都可以有方法
*/

package main

import "fmt"


type rect struct {
	width, height int
}

//这里的 area 方法有一个接收器(receiver)类型 rect。
func (r *rect) area() int {
	return r.width * r.height
}

//可以为值类型或者指针类型的接收器定义方法。这里是一个 值类型接收器的例子。
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

//实例：非结构体类型方法
type IntVector []int

func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area()) //以指针(或值)为接收者的方法被调用时，接收者既能为值又能为指针
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim()) // 以指针(或值)为接收者的方法被调用时，接收者既能为值又能为指针

	////实例：非结构体类型方法
	fmt.Println(IntVector{1, 2, 3}.Sum()) //  fmt.Println((&IntVector{1, 2, 3}).Sum())

	s := &IntVector{1, 2, 3, 4, 5, 6}
	fmt.Println(s.Sum())

}
