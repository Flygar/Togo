package main

import "fmt"

func main() {
	testData:=[]int{3,2,2,3}
	fmt.Println(removeElement2(testData,3))
	
}

func removeElement(nums []int, val int) int {
	c:=0
	for _,v :=range nums{
		if v==val{
			c+=1
		}
	}
	return len(nums)-c
}
func removeElement2(nums []int, val int) int {
	count := 0;
	pos := 0;
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			count++
			continue
		}
		nums[pos] = nums[i]
		pos++
	}
	fmt.Println(nums)
	return len(nums) - count

}