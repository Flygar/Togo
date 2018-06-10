package main

import (
	"fmt"
	"unsafe"
)

/*
定义一个 Rectangle 结构体，它的长和宽是 int 类型，并定义方法 Area() 和 Primeter()，然后进行测试。
*/

type rectangle struct {
	l,w int
}

func main() {
	r1 :=&rectangle{5,6}
	fmt.Println(Area(r1))
	fmt.Println(Primeter(r1))
	size := unsafe.Sizeof(r1)
	fmt.Println(size)

}
func Area(a *rectangle)(int)  {
	return a.l*a.w
}
func Primeter(a *rectangle)(int){
	return 2*(a.l+a.w)
}
