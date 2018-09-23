/*
结构体
*/

package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	//通过new来初始化
	st := new(person)
	st.name = "flyagr"
	st.age = 8
	//使用这个语法创建新的结构体元素
	fmt.Println(person{"Bob", 20})
	//你可以在初始化一个结构体元素时指定字段名字。
	fmt.Println(person{name: "Alice", age: 30})
	//省略的字段将被初始化为零值。
	fmt.Println(person{name: "Fred"})
	//& 前缀生成一个结构体指针。
	fmt.Println(&person{name: "Ann", age: 40}) // 底层仍然会调用 new ()；表达式 new(Type) 和 &Type{} 是等价的
	//使用.来访问结构体字段。
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	//也可以对结构体指针使用. - 指针会被自动解引用。
	sp := &s
	fmt.Println(sp.age) // (*sp).age
	//结构体是可变(mutable)的。
	sp.age = 51
	fmt.Println(sp.age)
}
