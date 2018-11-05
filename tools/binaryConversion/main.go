//在 Go 当中 string 底层是用 byte 数组存的
package main

import (
	"fmt"
	"strconv"
)

func main() {

	//bin2oct(2)
	var v int64 = 8 //默认10进制
	s8 := strconv.FormatInt(v, 10)
	fmt.Printf("%v\n", s8)

	var sv = "19584c4d"; // 16 to 10
	fmt.Println(strconv.ParseInt(sv, 16, 32))

}

//二进制2十进制
func bin2oct(x int64) {

}

//十进制2二进制
func oct2bin(x int64) {
	fmt.Println(strconv.FormatInt(x, 2))
}

//十进制2八进制
func oct2eight(x int64) {
	fmt.Println(strconv.FormatInt(x, 8))
}

//十进制2十六进制
func oct2hex(x int64) {
	fmt.Println(strconv.FormatInt(x, 16))
}

