package main
/*
chapter_10
*/

import (
	"fmt"
	"strings"
)

/*
struct:值类型,可以通过new来创建
组成结构体类型的那些数据称为 字段（fields）。每个字段都有一个类型和一个名字；在一个结构体中，字段名字必须是唯一的
*/
//定义结构体
type Person struct {
	name string
	age  int
}

//定义函数(接受一个指针)将name转换为大写
//不用指针(去掉*)的话不会对结构体产生影响
func UpName(p *Person) {
	p.name = strings.ToUpper(p.name)
	fmt.Printf("%v\n", p)
}

func main() {
	fmt.Println("--------Struct-1--------")
	var p1 Person = Person{"flygar", 20}
	UpName(&p1)
	fmt.Println(p1)

	fmt.Println("--------Struct-2--------")
	p2 := new(Person) //p2是个指针
	p2.name = "hello"
	(*p2).age = 1 //解指针,相当于*&Person,相当于Person
	UpName(&*p2)  //相当于&Person
	fmt.Println(p2)

	fmt.Println("--------Struct-3--------")
	p3 := &Person{"world", 2}
	UpName(p3)
	fmt.Println(p3)

	//递归结构体(链表，二叉树)
	type Tree struct {
		le   *Tree
		data float64
		ri   *Tree
	}

	//结构体工厂

	//带标签的结构体

	//匿名字段和内嵌结构体
	//小练习：创建一个结构体，它有一个具名的 float 字段，2 个匿名字段，类型分别是 int 和 string。通过结构体字面量新建一个结构体实例并打印它的内容。
	type struct1 struct {
		a float32
		int
		string
	}

	s1 := new(struct1)
	s1.a = 18.8
	s1.int = 20
	s1.string = "hello"
	fmt.Println(s1)

	//非结构体方法(数组求和)
	fmt.Println("--------Method-1---------")
	fmt.Println(test2{1, 2, 3, 4}.Sum())

	//结构体方法
	fmt.Println("--------Method-2---------")
	m1 := new(test3)
	m1.b="Flygar"
	m1.a=20
	fmt.Println(m1.info())
	fmt.Println(m1.info2("playing games"))

	//函数和方法的区别
	//函数将变量作为参数：Function1(recv)
	//方法在变量上被调用：recv.Method1(),不要忘记 Method1 后边的括号 (),否则会引发编译器错误.


}
func test(a, b int) (c int) {
	c = a + b
	return

}

//非结构体方法，数组求和
type test2 []int

func (he test2) Sum() (s int) {
	for _, v := range he {
		s += v
	}
	return

}

//结构体方法
type test3 struct {
	a int
	b string
}

func (m1 test3) info() (s string)  {
	fmt.Printf("hello,my name is %v,and my age is %v",m1.b,m1.a)
	return s

}
func (m1 test3) info2(h string) (s string)  {
	fmt.Printf("My hobbit is %v,hello,my name is %v,and my age is %v",h,m1.b,m1.a)
	return s

}
