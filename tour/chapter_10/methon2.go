package main

import "fmt"

/*
非结构体类型方法
*/

//定义数组求和方法
type intVector []int

func (p intVector) addSum() (s int) {
	for _,v :=range p{
		s +=v
	}
	return

}

func main() {
	s0:=intVector{1,2,3,4,5}
	fmt.Println(s0.addSum())

	
}
