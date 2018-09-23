package main

import "fmt"

func main() {
	isOdd(7)


}

//isOdd 判断奇偶性
func isOdd(num int) {
	if num <= 0 {
		fmt.Println("please enter a num > 0")
	} else if num%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd") //奇
	}
}
