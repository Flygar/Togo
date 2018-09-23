package main

import "fmt"

func main() {
	str1 := []string{"fltower", "fltow", "flight"}
	a := longestCommonPrefix(str1)
	fmt.Println(a)

}

func longestCommonPrefix(strs []string) string {
	c := 0
	l := len(strs[0])
	for k, v := range strs {
		if len(v) <= l {
			l = len(v)
			c = k
		}
	}
	//fmt.Println(c,l)
	t := ""
	for i := 1; i <= l; i++ {
		s := 0
		//fmt.Println(strs[c][:i])
		for j := 0; j < len(strs); j++ {
			if strs[c][:i] == strs[j][:i] {
				s += 1
			}
		}
		if s == len(strs) {
			t = strs[c][:i]
		}
	}
	return t

}
