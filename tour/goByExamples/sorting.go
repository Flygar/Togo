/*
排序：Go 的 sort 包实现了内置和用户自定义数据类型的排序 功能。我们首先关注内置数据类型的排序。
原地更新：改变原本序列
*/
package main

import (
	"fmt"
	"sort"
)
// 为了在 Go 中使用自定义函数进行排序，我们需要一个对应的类型。这里我们创建一个为内置 `[]string` 类型的别名的`ByLength` 类型，
type ByLength []string
//我们在类型中实现了 sort.Interface 的 Len，Less 和 Swap 方法，这样我们就可以使用 sort 包的通用 Sort 方法了，Len 和 Swap 通常在各个类型中都差 不多，Less 将控制实际的自定义排序逻辑。在我们的例 子中，我们想按字符串长度增加的顺序来排序，所以这里 使用了 len(s[i]) 和 len(s[j])。
func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {

	// 排序方法是针对内置数据类型的；这里是一个字符串的例子。
	// 注意排序是原地更新的，所以他会改变给定的序列并且不返回一个新值。
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// 一个 `int` 排序的例子。
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// 我们也可以使用 `sort` 来检查一个序列是不是已经是排好序的。
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)

	//--------------------------------------
	//自定义排序：按照元素的长度
	//通过将原始的 `fruits` 切片转型成 `ByLength` 来实现我们的自定排序了。然后对这个转型的切片使用 `sort.Sort` 方法。
	//类似的，参照这个创建一个自定义类型的方法，实现这个类型的 这三个接口方法，然后在一个这个自定义类型的集合上调用 sort.Sort 方法，我们就可以使用任意的函数来排序 Go 切片了
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)

	
}
