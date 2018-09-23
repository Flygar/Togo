package main

import (
	"strings"
	"fmt"
)

func main() {
	s:="Hello  23"
	fmt.Println(lengthOfLastWord(s))
	
}

func lengthOfLastWord(s string) int {
	temp:=strings.Fields(s)
	//fmt.Println(temp[len(temp)-1])
	//fmt.Println(len(temp[len(temp)-1]))

	if len(temp)==1{
		return 0
	}
	return len(temp[len(temp)-1])

}