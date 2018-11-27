package main

import (
	"fmt"
	"sort"
)

// var slice1 []type = make([]type, len, cap) //左边的[]type可以省略，go根据右边做类型推断

//构造长度为1000的slice的切片，并且出去11，23整除的数字

func main() {

	//切片的声明和初始化
	var s = []string{1: "a", 100: "b"}
	s2 := [3]int{2, 5, 1} // 不能直接 `s2 := [3]int{2, 5, 1}[:]` 生成切片
	s4 := s2[:]
	s3 := []int{7, 3, 10}
	fmt.Println(s, s2, s3)

	//如果参数被存储在一个 slice 类型的变量 slice 中，则可以通过 slice... 的形式来传递参数，调用变参函数。
	fmt.Println(min(s4...))
	fmt.Println(min(s3...))

	//实例:生成相同的切片
	//make([]int, 50, 100) //数组长度为100，切片长度为50
	//new([100]int)[0:50]  //先通过new创建长度为100的数组，然后截取下标0-50的切片

	//---------------------------------------------

	//追加
	a := []int{1, 2, 3, 4, 5}
	b := []int{10, 9, 8, 7}

	//将元素 x , x2 追加到切片 b：
	b = append(b, 6, 6)
	fmt.Println(b)

	//将切片 b 的元素追加到切片 a 之后：
	a = append(a, b...)
	fmt.Println(a)

	//删除位于索引 i 的元素：
	a = append(a[:i], a[i+1:]...)

	//删除切片 a 中从索引 i 至 j 位置的元素：
	a = append(a[:i], a[j:]...)

	//为切片 a 扩展 j 个元素长度：
	a = append(a, make([]T, j)...)

	//在索引 i 的位置插入元素 x：
	a = append(a[:i], append([]T{x}, a[i:]...)...)

	//在索引 i 的位置插入长度为 j 的新切片：
	a = append(a[:i], append(make([]T, j), a[i:]...)...)

	//在索引 i 的位置插入切片 b 的所有元素：
	a = append(a[:i], append(b, a[i:]...)...)

	//取出位于切片 a 最末尾的元素 x：
	x, a = a[len(a)-1], a[:len(a)-1]

	//----------------------------------------------

	//复制
	//复制切片 a 的元素到新的切片 c 上
	c := make([]int, len(a))
	copy(c, a) // 副本
	fmt.Println(c)

	//引用传递
	fmt.Println(a)
	d := a // 引用传递，指针向 a 的地址; 对a或d的修改其实是修改同一个地址上的值
	fmt.Println(d)
	a[0] = 1111111
	fmt.Println(d)
	d[1] = 22222222
	fmt.Println(a)

	//-----------------------------------------------

	//排序
	fmt.Println("b: ", b)
	sort.Ints(b)
	fmt.Println("排序后的b:", b)
	newA := []string{"cab", "abc", "bca", "aab"}
	fmt.Println("a: ", newA)
	sort.Strings(newA) // 先按第一个字母，再按第二个字母
	fmt.Println("排序后的a:", newA)

}

//min...求slice中的最小元素的值
func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}
