package main

import "fmt"

/*
计算正方形和圆的面积
*/

type areaSize interface {
	area() float64
}

func showAreaSize(a *areaSize) {
	fmt.Printf("大小是： %v\n",(*a).area())
}

type square2 struct {
	a float64
}

func (s *square2) area() float64 {
	return s.a*s.a
}

type circle2 struct {
	r float64
}

func (c *circle2) area()  float64{
	return 3.14*c.r*c.r
}

func main() {
	s0 :=&square2{11}
	c0 :=&circle2{11}

	v0 :=[]areaSize{s0,c0}

	if t, ok := v0[0].(*square2); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}

	if t, ok := v0[0].(*circle2); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}


	for k,_ :=range v0{
		showAreaSize(&v0[k])
	}


	switch t := v0[0].(type) {
	case *square2:
		fmt.Printf("Type Square %T with value %v\n", t, t)
	case *circle2:
		fmt.Printf("Type Circle %T with value %v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}

}
