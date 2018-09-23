package main

import "fmt"

// 只包含一个方法的接口用[e]r结尾
type shaper interface {
	area() float64
}

//结构体
type square struct {
	side float64
}

//结构体square的方法
func (s *square) area() float64 {
	return s.side*s.side
}

func main() {
	s0 := &square{3}

	//var areaIntf shaper
	//areaIntf = s0
	// 简短写法
	//areaIntf := Shaper(sq1)
	// 更简短写法
	areaIntf :=s0

	//通过结构体方法实现
	fmt.Printf("%f\n",s0.area())
	//通过接口实现
	fmt.Println(areaIntf.area())

	
}
