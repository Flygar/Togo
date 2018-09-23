package main

import (
	"fmt"
	"net/http"
)


func main() {
	climbStairs(2)
	
}

func climbStairs(n int) int {
	http.Server{}
	a:=[]int{1,2}
	s:=[]int{1,2}
	temp:=make([]int,0)
	//var s []int
	for i:=0;i<n;i++{
		fmt.Printf("目前a切片的元素有：%v | 准备执行各元素+1，+2的操作\n",a)
		fmt.Println("-----开始执行------")
		//脑容量不够大就不要在循环中尝试改变变量k,v,a的值
		for _,v :=range a{
			fmt.Println("小循环中a的值：如下")
			fmt.Println(a)
			temp=append(temp,v+1,v+2)
			s=append(s,temp...)
			fmt.Println("+1+2操作后，小循环中s的值：如下")
			fmt.Println(s)
		}

		a=s-a
		fmt.Println("-----结束执行------")
		fmt.Println("把s的值赋给a，s的值：",s)

	}
return 0
}