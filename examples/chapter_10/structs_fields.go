package main

import (
	"fmt"
	"unsafe"
	"reflect"
)

type t struct {
	i1  int "tag1"
	f1  float32 "tag2"
	str string "tag3"
}

func main() {

	fmt.Println("-0-")
	//分配内存并零值化
	var ms3 t
	ms3 =t{0,0.0,"Hi"}
	//ms3.str = "hello"
	fmt.Printf("%v", ms3.str)
	//查看结构体实例占用多少内存
	size3 := unsafe.Sizeof(ms3)
	fmt.Println(size3)
	//查看tag索引
	fmt.Println(reflect.TypeOf(t{}).Field(1).Tag)

	fmt.Println("\n-1-")
	ms0 := t{0, 0.0, "hello"}
	fmt.Println(ms0)

	fmt.Println("-2-")
	//使用new()函数给一个指针变量分配内存，指向t的地址;var t *T = new(T)
	//var ms1 *t = new(t)
	ms1 := new(t)
	//如果定义是包范围的，但是分配却没有必要在开始就做
	//var ms1 *t
	//ms1 = new(t)
	ms1.i1 = 1
	//解指针,反向引用
	(*ms1).f1 = 1.1
	ms1.str = "world"
	fmt.Printf("The int is: %d\n", ms1.i1)
	fmt.Printf("The float is: %f\n", ms1.f1)
	fmt.Printf("The string is: %s\n", ms1.str)
	fmt.Println(ms1)

	fmt.Println("-3-")
	//&t{a, b, c} 是一种简写，底层仍然会调用 new ()
	ms2 := &t{2, 2.2, "flygar"}
	fmt.Printf("int:%v", ms2.i1)
	fmt.Println("\n", ms2)

}
