package main

/*
注意：
map 无序排列
map 引用类型
map 非线程安全 当并行访问一个共享的 map 类型的数据，map 数据将会出错
*/

import (
	"fmt"
	"sort"
)

func main() {

	//定义map
	var mapLit map[string]int
	//初始化
	//mapLit = map[string]int{}
	//初始化并赋值
	mapLit = map[string]int{"one": 1, "two": 2}
	mapLit["three"] = 3
	fmt.Println(mapLit)

	//定义map并初始化
	m := map[string]string{}
	m["key3"] = "val3"
	//定义map并初始化并赋值
	m = map[string]string{"key1": "hello", "key2": "world"} //不会输出m["key3"]，因为被初始化了
	fmt.Println(m)

	//通过make()定义map并初始化
	var map1 = make(map[string]int)
	map1["map3"] = 3
	fmt.Println(map1)

	//下面2者等价
	mapCreated1 := make(map[string]float32)
	mapCreated2 := map[string]float32{}
	fmt.Println(mapCreated1, mapCreated2)

	//用切片作为map的value
	mp1 := make(map[int][]int)
	//mp2 := make(map[int]*[]int)
	mp1[1] = []int{1, 2, 3}
	mp1[2] = []int{13, 22, 31}
	mp1[3] = []int{}

	//判断，如果key不存在，则创建，默认值为666
	key1 := int(1)
	if val1, ok := mp1[key1]; !ok {
		mp1[key1] = []int{666}
		fmt.Println(mp1)
	} else {
		fmt.Println(val1)
		delete(mp1, key1)
		fmt.Println(mp1)
	}
	fmt.Println(mp1)

	//练习：创建一个 map 来保存每周 7 天的名字，将它们打印出来并且测试是否存在 "星期二" 和 "星期八"。
	dateMap := make(map[string]string)
	dateMap["星期一"] = "Monday"
	dateMap["星期二"] = "Tuesday"
	dateMap["星期三"] = "Wednesday"
	dateMap["星期四"] = "Thursday"
	dateMap["星期五"] = "Friday"
	dateMap["星期六"] = "Saturday"
	dateMap["星期天"] = "Sunday"
	for k, v := range dateMap {
		fmt.Printf("%s | %s\n", k, v)
	}
	key := "星期二"
	if value, ok := dateMap[key]; ok {
		fmt.Println("存在：", value)
	}

	//Map类型的切片
	//Version A
	//第一次make()分配切片，第二次make()分配切片中每个map元素
	items := make([]map[int]int, 2)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 1
	}
	fmt.Printf("Version A: Value of items: %v\n", items)

	// Version B: NOT GOOD!
	items2 := make([]map[int]int, 5)
	for _, item := range items2 { // item是切片元素的一个拷贝，并不是(指向)items2中的元素(地址)
		item = make(map[int]int, 1)
		item[1] = 2
	}
	fmt.Printf("Version B: Value of items: %v\n", items2)

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

	//Map键值对调
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
}
