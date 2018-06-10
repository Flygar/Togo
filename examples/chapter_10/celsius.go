package main

import (
	"fmt"
	"strconv"
)

/*
为 float64 定义一个别名类型 Celsius，并给它定义 String()，它输出一个十进制数和 °C 表示的温度值。
*/

type Celsius float64

func (c Celsius) String() string{
	return strconv.FormatFloat(float64(c), 'f', 1, 32) + " °C"

}
func main() {
	c := Celsius(2.1)
	fmt.Println(c)

}
