package main

import "fmt"

func main() {

	a := romanToInterger("IVI")

	//b := "hello world"
	//fmt.Println(string(b[0]))
	//for k, v := range b {
	//	fmt.Println(k, v)
	//}
	fmt.Println(a)

}

func romanToInterger(s string) int {
	dict := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

	ret := 0
	//s=
	for k, v := range s {
		if k == len(s)-1 {
			ret += dict[string(v)]
		} else {
			if dict[string(v)] < dict[string(s[k+1])] {
				ret -= dict[string(v)]
			} else {
				ret += dict[string(v)]
			}
		}
		//fmt.Println(k,"+",string(v))
		//ret += dict[string(v)]
	}

	return ret
}
