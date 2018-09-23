package main

import "fmt"

func main() {
	nums := []int{4,7,9,100}
	a := removeDuplicates(nums)
	fmt.Println(a)

}

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	var length int
	for k, v := range nums {
		if nums[k] != nums[0] {
			nums[0] = v
			length++
		}
	}
	return length+1
}
