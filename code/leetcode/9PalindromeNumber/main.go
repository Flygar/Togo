package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 65
	//返回a所代表的ASCII编码的字符：A
	fmt.Printf("%v\n", string(a))

	//返回a所代表的ASCII编码的字符：A 的十进制表示：65
	fmt.Printf("%v\n", []byte(string(a)))

	//转成字符串形式
	fmt.Printf("%v\n", string([]byte(string(a))))

	//转成字符串形式
	fmt.Printf("%v", strconv.Itoa(a))

}

func Palindrome(num int) bool {
	str1 := strconv.Itoa(num)
	for i := range str1 {
		if str1[i] != str1[len(str1)-1-i] {
			return false
		}
	}
	return true
}
