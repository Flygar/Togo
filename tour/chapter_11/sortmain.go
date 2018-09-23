package main

import (
	"./sort"
	"fmt"
)
func inits()  {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	b :=data
	a := sort.Test3(data)
	//b:=sort.Sorter(a)
	//sort.Sort(b)
	//系统自动进行转换，可以直接下面最简短的形式
	sort.Sort(a)
	if !sort.IsSorted(a) {
		panic("fails")
	}
	fmt.Printf("The sorted array is: %v\n", a)
	fmt.Println(data)
	fmt.Println(b)

	//fmt.Println(b)
}





func main() {
	inits()
	
}
