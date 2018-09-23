package main

import "fmt"

/*
为 int 定义一个别名类型 Day，定义一个字符串数组它包含一周七天的名字，为类型 Day 定义 String() 方法，它输出星期几的名字。使用 iota 定义一个枚举常量用于表示一周的中每天（MO、TU...）。
*/

type Day int

const (
	MO Day = iota
	TU
	We
	TH
	FR
	SA
	SU
)

var dayName = []string{"Monday", "tuesday", "wednesday", "thursday", "Friday", "Saturday", "Sunday"}

func (d Day) String() string {
	return dayName[d]

}

func main() {
	var th Day = 3
	fmt.Println(th)
	var day Day= SU
	fmt.Println(day)
	fmt.Println(0, MO, 1, TU)
	fmt.Println(dayName[MO])


}
