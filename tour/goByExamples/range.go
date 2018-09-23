package main

/*
数组和切片中：k是下标，v是元素的值
map中：k是键(不是下标)，v是值
*/
import "fmt"

func main() {

	// range 的拷贝
	nums := []int{2, 3, 4}
	for _, num := range nums {  //  num只是nums中元素的一个拷贝，所以下面的 num=num+1 并不会对原切片 nums 产生修改
		num=num+1
	}
	fmt.Println(nums)

	// range 修改原切片
	for k,v := range nums{
		nums[k]=v+1  // 操作的是原切片
	}
	fmt.Println(nums)

	// map
	map1 := make(map[string]string)
	//map1 := map[string]string{"a": "apple", "b": "banana"}
	map1["name"]="Flygar"
	map1["age"]="8"
	map1["address"]="Earth"
	for k,v :=range map1{
		fmt.Println(k,v)
	}

	//range 在字符串中迭代 unicode 码点(code point)。 第一个返回值是字符的起始字节位置，然后第二个是字符本身。
	for i, c := range "go我和你" {
		fmt.Println(i, c)
	}
}
