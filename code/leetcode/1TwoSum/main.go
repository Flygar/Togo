package main

import "fmt"

func main() {
	//slice1 := make([]int, 0)
	slice1 := []int{2, 7, 11}
	//向slice追加元素
	slice1 = append(slice1, 15)
	target := 9
	ret := fun1(&slice1, target)
	fmt.Println(ret)
}

func fun1(nums *[]int, value int) []int {
	map1 := make(map[int]int)
	//记得比对用指针的运行时间
	for id, num := range *nums {
		map1[value-num] = id
	}
	fmt.Println(map1)

	for id, num := range *nums {
		fmt.Printf("Checking index %v=%v\n", id, num)
		if val, ok := map1[num]; ok && id != val {
			return []int{id, val}
		}
	}

	return nil
}


//循环遍历的方法
func twoSum(nums []int, target int) []int {
	indexes := make([]int, 0)
	for i, v := range nums {
		for in, va := range nums {
			if i != in && v+va == target && len(indexes) < 2 {
				indexes = append(indexes, i, in)
			}
		}
	}
	return indexes
}
