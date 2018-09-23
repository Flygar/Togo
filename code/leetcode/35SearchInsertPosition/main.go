package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a := []int{1, 3, 5, 6}
	target := 2
	ret := searchInsert(a, target)
	fmt.Println(ret)
	ret2 := f2(a, target)
	fmt.Println(ret2)

	//b := []string{"a", "b", "c"}
	//c := strings.Join(b, "-")
	//fmt.Printf("%v,%T", b, c)

}

func searchInsert(nums []int, target int) int {
	for k, v := range nums {
		if v == target {
			return k
		} else if v > target {
			return k

		}
	}
	return len(nums)
}

func f2(nums []int, target int) int {
	b := ""
	for _, v := range nums {
		b += strconv.Itoa(v)
	}
	fmt.Println(b)

	if strings.Contains(b, strconv.Itoa(target)) {
		fmt.Println("hhh")
		return strings.Index(b, strconv.Itoa(target))

	} else {
		fmt.Println("----")
		for k, v := range b {
			test, _ := strconv.Atoi(string(v))
			//fmt.Printf("%v,%T\n",test,test)
			if test > target {
				//fmt.Println(test,target)
				return k
			}
		}

	}

	return -1

}

//二分查找实现
