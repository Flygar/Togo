package main

import (
	"fmt"
	"strconv"
)

func main() {

	a := []int{4,3,2,9}
	//a := fmt.Sprint(test)
	fmt.Println(plusOne(a))
}

func plusOne(digits []int) []int {
	s := ""
	for _, v := range digits {
		s += strconv.Itoa(v)
	}
	t, _ := strconv.Atoi(s)
	t += 1

	b := strconv.Itoa(t)
	//fmt.Printf("%T",b)
	//fmt.Println(b[1])

	a := make([]int, len(digits))
	for k, v := range b {
		//fmt.Println(k,v)
		a[k],_=strconv.Atoi(string(v))
	}

	return a
}
