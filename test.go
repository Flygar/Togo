package main

import "fmt"

func main() {
	ages := []int{1, 2, 3, 5}
	b := ages
	c :=ages
	//c[0]++
	b[0]++
	fmt.Println(ages)
	fmt.Println(b)
	fmt.Println(c)
}