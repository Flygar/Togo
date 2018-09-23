package main
/*
编写一个程序，从键盘读取输入。当用户输入 'S' 的时候表示输入结束，这时程序输出 3 个数字：
i) 输入的字符的个数，包括空格，但不包括 '\r' 和 '\n'
ii) 输入的单词的个数
iii) 输入的行数
*/

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)
var c,w,l int

func main()  {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input, type S to stop:")
	input, err := inputReader.ReadString('S')

	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	fmt.Printf("Your name is %s", input)
	// For Unix: test with delimiter "\n", for Windows: test with "\r\n"
	switch input {
	case "Philip\n":  fmt.Println("Welcome Philip!")
	case "Chris\n":   fmt.Println("Welcome Chris!")
	case "Ivo\n":     fmt.Println("Welcome Ivo!")
	default: fmt.Printf("You are not welcome here! Goodbye!")
	}

	// version 2:
	switch input {
	case "Philip\n":  fallthrough
	case "Ivo\n":     fallthrough
	case "Chris\n":   fmt.Printf("Welcome %s\n", input)
	default: fmt.Printf("You are not welcome here! Goodbye!\n")
	}

	// version 3:
	switch input {
	case "Philip\n", "Ivo\r\n":   fmt.Printf("Welcome %s\n", input)
	default: fmt.Printf("You are not welcome here! Goodbye!\n")
	}
}

func Counters(input string) {
	c += len(input) - 1 // -2 for \n
	w += len(strings.Fields(input))
	l++
}