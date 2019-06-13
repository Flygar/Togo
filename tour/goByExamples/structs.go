/*
结构体:
注意1: 值类型，声明的时候默认零值;数组类型的默认零值和他的元素有关;指针，slice和map的零值都是nil，即还没分配空间
注意2: struct每个field上可以写上一个tag，该tag可以通过反射机制获取，常见的使用场景就是序列化和反序列化。
序列化：可以把一个变量序列化成字符串再返回去，可以的通用格式为JSON.(将struct变量进行json处理后给其他程序，如php或jquery处理，但那边不适合习惯大写，这时候就需要用到tag)

*/

package main

import "fmt"

type Person struct {
	Name string  ‘name’// field
	age  int     // field
	scores [5]float64
	ptr *int             // 引用类型，默认值nil，还没有分配空间。 使用前需要new
	slice []int          // 引用类型，默认值nil，还没有分配空间。 使用前需要make
	map1 map[string][string]   // 引用类型，默认值nil，还没有分配空间。 使用前需要make
}


//结构体中结构体
type Point struct {
	x,y int
}
type Rect struct {
	leftUp,rightDown Point
//r1:=Rect{Point{1,2},Point{3,4}}

//tag及序列化的介绍
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	//通过new来初始化,是个指向person的指针。这里注意和默认的值类型拷贝做区别。
	st := new(person) // var st *person = new(person)
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



	//通过tag，序列化为小写，保持兼容性
	//1. 创建一个Monster变量
	monster := new(Monster)
	monster.Name = "牛魔王"
	monster.Age = 500
	monster.Skill = "芭蕉扇"
	//2. 将monster变量序列化为JSON格式字串. 返回小写的fields，保持兼容性
	jsonMonster, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("序列化json时错误：", err)
	}
	fmt.Println("json:", string(jsonMonster))


}
