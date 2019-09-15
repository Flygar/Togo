package main

/*
注意0： nil标志符用于表示interface、函数、maps、slices和channels的“零值”。字符串的0值是""
注意1:map：需要声明并初始化，长度可以动态变化(map,slice都能动态增加，slice of map)，无序排列
注意2:slice,map还有function不可以，因为这几个没法用 == 来判断
注意3:声明是不会分配内存的，初始化需要make，分配内存后才能赋值和使用
注意4:map的遍历不能用for循环，得用for range。
注意5: cap: 返回的是数组切片分配的空间大小, 根本不能用于map。make: 用于slice，map，和channel的初始化
*/
import (
	"fmt"
	"sort"
)

func main() {
	//make声明并初始化
	//map1 := map[string]string{"a": "apple", "b": "banana"}
	map1 := make(map[string]string) ////make：声明并初始化（赋值）
	map1["name"] = "Flygar"
	map1["age"] = "08"
	map1["address"] = "Earth"
	fmt.Println(map1)

	var map2 map[string]int                   // 声明
	map2 = map[string]int{"one": 1, "two": 2} // 初始化（赋值）
	map2["age"] = 8
	fmt.Println(map2)
	
	//支持value是个函数(工厂模式)
	m1 := make(map[int]func(op int) int)
	m1[1] = func(op int) int { return op }
	m1[2] = func(op int) int { return op*2 }
	m1[3] = func(op int) int { return op*3 }
	t.Log(m1[1](2), m1[2](2), m1[3](2))

	//用切片slice作为 map 的value
	map3 := make(map[int][]float64)
	map3[1] = []float64{1.1, 1.2, 1.3}
	fmt.Println(map3)

	//slice of map ，map切片
	//slice of map,动态增加map
	monster := make([]map[string]string, 2)
	if monster[0] == nil {
		monster[0] = make(map[string]string)
		monster[0]["name"] = "monster niu"
		monster[0]["age"] = "500"
		monster[0]["wife"] = "monster TieShan"
	}

	if monster[1] == nil {
		monster[1] = make(map[string]string)
		monster[1]["name"] = "monster rabbit"
		monster[1]["age"] = "400"
	}
	//下面这个错误
	//if monster[2]==nil {
	//	monster[2]=make(map[string]string)
	//	monster[2]["name"]="monster red child"
	//	monster[2]["age"]="300"
	//}
	
	monsterNew := make(map[string]string, 2)
	monsterNew["name"] = "火云邪神"
	monsterNew["age"] = "200"

	monster=append(monster,monsterNew)
	fmt.Println(monster)

	//用map作为map的value，需要二次make
	//例子：key是学号，value是个map存放“姓名”和“性别”
	studentMap := make(map[int]map[string]string, 10)
	studentMap[101] = make(map[string]string)
	studentMap[101]["姓名"] = "宋江"
	studentMap[101]["性别"] = "男"
	studentMap[101]["称谓"] = "及时雨"
	studentMap[101]["Address"] = "黄浦江"

	studentMap[102] = make(map[string]string)
	studentMap[102]["姓名"] = "吴用"
	studentMap[102]["性别"] = "男"
	studentMap[102]["称谓"] = "智多星"

	studentMap[103] = make(map[string]string)
	studentMap[103]["姓名"] = "卢俊义"
	studentMap[103]["性别"] = "男"
	studentMap[103]["称谓"] = "玉麒麟"

	//fmt.Println(studentMap)

	for k, v := range studentMap {
		//fmt.Println(k,v)
		fmt.Println("##########学号：", k)
		for k2, v2 := range v {
			//fmt.Println(k2)
			fmt.Printf("%v: %v\n", k2, v2)
		}
	}
}

	//查找判断key值是否存在,存在则ok为true
	if _, ok := map1["name"]; ok {
		fmt.Println("exist")
	}

	//删除key值
	delete(map1, "name") // remove map1[key1]
	fmt.Println(map1)
	//----------------------------------------------
	
	//修改map中元素的值
	type data struct {  
		name string
	}
	m := map[string]data{"x": {name:"one"}, "y": {name: "two"}}
	//m["y"].name = "three" // cannot assign to m["y"].name
	tempVarA:=m["y"]
	tempVarA.name="three"
	//不推荐下面这个修改方法，元素多了难写
	// m["y"]=data{name:"three"}
	// fmt.Println(m["y"].name)

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
