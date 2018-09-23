package main

import "fmt"

func main() {
	nums:=[]int{-1,3,9,-3,-6,5,8,-3}
	fmt.Println(maxSubArray(nums))
	
}

func maxSubArray(nums []int) int {
	if nums==nil{
		return 0
	}
	//temp :=make([]int,len(nums))

	//temp:=func(x []int) []int{
	//	temp:=make([]int,0)
	//	for k,v :=range x{
	//		if k<len(x)-1{
	//			temp=append(temp,v+x[k+1])
	//		}
	//	}
	//	return temp
	//}(nums)

	total:=0
	indexS:=0
	indexE:=0
	for i:=0;i<len(nums);i++{
		for j:=i;j<len(nums);j++{
			//temp:=nums[i:j]
			s:=0
			for _,v :=range nums[i:j]{
				s+=v
			}
			if s>total{
				total=s
				indexS=i
				indexE=j
			}
		}
	}
	fmt.Println(nums[indexS:indexE])
	return total

}