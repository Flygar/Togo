package main

import (
	"fmt"
	"strings"
)

func main() {
	haystack := "aaaaa"
	needle := "bba"
	fmt.Println(strStr(haystack, needle))
}
func strStr(haystack string, needle string) int {
	if strings.Contains(haystack, needle) {
		fmt.Println("-----")
		return strings.Index(haystack, needle)
	}
	return -1
}
