package main

import (
	"fmt"
)

func main() {
	var a byte
	//a='A'
	a = 65
	var b uint8
	b = 'B'
	//b=66
	fmt.Printf("a的类型: %T, a的值: %v, a的ASCII符号:%v"+"\n"+"b的类型: %T, b的值: %v, a的ASCII符号:%v", a, a, string(a), b, b, string(b))

}

func isValid(s string) bool {
	var stack []rune
	paren := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, v := range s {
		if _, present := paren[v]; !present {
			stack = append(stack, v)
			fmt.Println(present, " - ", v, " - ", stack)

		} else {
			return false
		}
	}

	return len(stack) == 0

}
