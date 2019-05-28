package main

import "fmt"

func main() {
	//fcRange("hello")
	//testFor()
	//testGoto()
	//fcLabel()
	//fcLabel2()

}

//可以迭代任何一个集合（包括数组和 map)
func fcRange(s string) {
	for k, v := range s {
		fmt.Println(k, string(v),v)
	}
} 

//testFor golang for
func testFor() {
	for i := 0; i <= 5; i++ {
		fmt.Println("Flygar,i= ", i)
		for j := 0; j <= 10; j++ {
			if j == 3 {
				//break  //跳出本层循环
				continue //进入下次循环
			}
			fmt.Println("j= ", j)
		}
	}
}

//testGoto...不推荐使用 LABEL 和 goto
func testGoto() {
	i := 0
START:
	fmt.Printf("The counter is at %d\n", i)
	i++
	if i < 15 {
		goto START
	} else {
		fmt.Println("goto over")
	}

	fmt.Println("The End")

}

//fcLabel LABEL标签
func fcLabel() {
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
}

//fcLabel2 LABEl标签
func fcLabel2() {
	i := 0
HERE:
	fmt.Println(i)
	i++
	if i == 5 {
		return
	}
	goto HERE
}
