package main

import "fmt"

/*
匿名结构体
创建一个结构体，它有一个具名的 float 字段，2 个匿名字段，类型分别是 int 和 string。通过结构体字面量新建一个结构体实例并打印它的内容。
*/

//匿名
type s0 struct {
	int
	string
}

//内嵌
type s1 struct{
	a float32
	s0
}

func main() {

	t2 := &s1{0.0,s0{0,"Hi"}}
	fmt.Println(t2.s0.string)

}
