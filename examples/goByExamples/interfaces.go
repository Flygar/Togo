/*
接口(Interfaces): 方法签名(signatures)的集合;但是这些方法不包含（实现）代码：它们没有被实现（它们是抽象的）。接口里也不能包含变量。
空接口:可以存储任意类型的数值
接口类型的变量可以保存任何实现了这些方法的值
接口感觉有点多态的意思
*/

package main

import (
	"fmt"
	"math"
)

// 这里是一个几何体的基本接口。
type geometry interface {
	area() float64
	perim() float64
}

// 定义结构体类型 `rect2` 和 `circle` 准备实现接口
type rect2 struct {
	width, height float64
}
type circle struct {
	radius float64
}

// 要在 Go 中实现一个接口，我们就需要实现接口中的所有方法。这里我们在 `rect` 上实现了 `geometry` 接口。
func (r rect2) area() float64 {
	return r.width * r.height
}
func (r rect2) perim() float64 {
	return 2*r.width + 2*r.height
}

// 在 `circle` 上实现了 `geometry` 接口。
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 如果一个变量具有接口类型，那么我们可以调用指定接口中的方法。这里有一个通用的 `measure` 函数，利用它来在任何的 `geometry` 上工作。
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := &rect2{width: 3, height: 4} // 接口中的方法是值类型而不是指针
	c := &circle{radius: 5}          // 接口中的方法是值类型而不是指针
	// 结构体类型 `circle` 和 `rect` 都实现了 `geometry`接口，所以我们可以使用它们的实例作为 `measure` 的参数。
	measure(r)
	measure(c)

	//结构体 rect2 实现了接口 geometry；所以可以将一个 rect2 类型的变量赋值给一个接口类型的变量：geo = r
	fmt.Println("------------------")
	var geo geometry
	geo = r
	fmt.Println(geo.area())
	fmt.Println(geo.perim())
	fmt.Println("------------------")

	//循环输出
	shapes := []geometry{r, c}
	fmt.Println("Looping through shapes for area and perim...")
	for k := range shapes {
		fmt.Println("Shape details: ", shapes[k])
		fmt.Println("Area of this shape is: ", shapes[k].area())
		fmt.Println("Perim of this shape is: ", shapes[k].perim())

	}

	// 空接口存储任意类型
	var a interface{}
	a = "hello"
	a = 25
	fmt.Println(a)

}
