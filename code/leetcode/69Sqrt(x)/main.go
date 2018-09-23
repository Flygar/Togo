package main

import (
	"math"
	"fmt"
)

// 啊，我放弃思考。有标准库干深墨不用呢

func main() {
	a:=99
	fmt.Println(mySqrt(a))
	
}

func mySqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}