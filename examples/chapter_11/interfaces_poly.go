package main

import "fmt"

type clacArea interface {
	area() int
}

type test1 struct {
	side int
}

func (t *test1) area() int {
	return t.side * t.side
}

type test2 struct {
	length, width int
}

func (t *test2) area() int {
	return t.length * t.width

}

func main() {
	t1 := &test1{2}
	t2 := &test2{2, 3}
	//t3:=clacArea(t1)
	//t4:=clacArea(t2)
	t3 := []clacArea{t1, t2}
	//fmt.Println(t3[0].area())
	//fmt.Println(t3[1].area())
	for k, _ := range t3 {
		fmt.Printf("%v | ", t3[k])
		fmt.Printf("%v \n", t3[k].area())
	}

}
