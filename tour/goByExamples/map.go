package main

/*
map：需要声明并初始化
*/
import (
	"fmt"
	"sort"
)

func main() {
	//make声明并初始化
	//map1 := map[string]string{"a": "apple", "b": "banana"}
	map1 := make(map[string]string)
	map1["name"] = "Flygar"
	map1["age"] = "08"
	map1["address"] = "Earth"
	fmt.Println(map1)

	var map2 map[string]int                   // 声明
	map2 = map[string]int{"one": 1, "two": 2} // 初始化
	map2["age"] = 8
	fmt.Println(map2)

	//用切片作为 map 的value
	map3 := make(map[int][]float64)
	map3[1] = []float64{1.1, 1.2, 1.3}
	fmt.Println(map3)

	//判断key值是否存在,存在则ok为true
	if _, ok := map1["name"]; ok {
		fmt.Println("exist")
	}

	//删除key值
	delete(map1, "name") // remove map1[key1]
	fmt.Println(map1)

	//----------------------------------------------

	//map的排序
	//需要将 key（或者 value）拷贝到一个切片，再对切片排序（使用 sort 包），然后可以使用切片的 for-range 方法打印出所有的 key 和 value
	//Map的排序
	var (
		barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
			"delta": 87, "echo": 56, "foxtrot": 12,
			"golf": 34, "hotel": 16, "indio": 87,
			"juliet": 65, "kili": 43, "lima": 98}
	)
	fmt.Println("--------Map排序--------")
	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}
	keys := make([]string, len(barVal))
	i := 0
	for k := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println("sorted:")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v / ", k, barVal[k])
	}

	//Map键值对调;指定了 map 的长度 (len(barVal))，map 是可以动态增长的。
	fmt.Println("\n--------Map键值对调-------")
	invMap := make(map[int]string, len(barVal))
	for k, v := range barVal {
		invMap[v] = k
	}
	fmt.Println("inverted:")
	for k, v := range invMap {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}
	//排序
	fmt.Println("\n---------Map排序--------")
	key2 := make([]int, len(invMap))
	i = 0
	for k := range invMap {
		key2[i] = k
		i++
	}
	sort.Ints(key2)
	for _, v := range key2 {
		fmt.Println(invMap[v], v)
	}


	//----------------------------------------------

	//map类型的切片
	//应当像 A 版本那样通过索引使用切片的 map 元素。在 B 版本中获得的项只是 map 值的一个拷贝而已，所以真正的 map 元素没有得到初始化。
	// Version A:
	items := make([]map[int]int, 5)
	for i:= range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)

	// Version B: NOT GOOD!
	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int, 1) // item is only a copy of the slice element.
		item[1] = 2 // This 'item' will be lost on the next iteration.
	}
	fmt.Printf("Version B: Value of items: %v\n", items2)




}
