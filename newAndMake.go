package main
/*
可以使用 make() 的三种类型：slices/maps/channels
可以使用 new() 的情况：值类型;返回一个指针，指向内存地址
map()是引用类型
struct 值类型;可以用new()创建

*/

type Foo map[string]string
type Bar struct {
	thingOne string
	thingTwo int
}
func main() {
	// OK
	y := new(Bar)
	(*y).thingOne = "hello"
	(*y).thingTwo = 1

	// NOT OK 试图 make() 一个结构体变量，会引发一个编译错误
	z := make(Bar) // 编译错误：cannot make type Bar
	(*z).thingOne = "hello"
	(*z).thingTwo = 1

	// OK
	x := make(Foo)
	x["x"] = "goodbye"
	x["y"] = "world"

	// NOT OK map()引用类型，不能用new()
	u := new(Foo)
	(*u)["x"] = "goodbye" // 运行时错误!! panic: assignment to entry in nil map
	(*u)["y"] = "world"
	
}
