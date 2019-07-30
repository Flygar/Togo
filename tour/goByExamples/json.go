/*
json编解码
*/
package main

import (
	"encoding/json"
	"fmt"
)

//自定义类型的编码和解码。
type Response1 struct {
	Page   int
	Fruits []string
}
type Response2 struct {
	Page   int      `json:"page"`   
	Fruits []string `json:"-"`      
}

type Dummy struct {
	// Name    string  `json:"name"`   //解码json时，只取想要的字段（不一定json中的字段你struct中都要定义）
	Number  int64   `json:"-"`         //解码json时，即使这个field有值，忽略并设置它的value为相应类型的零值
	Pointer *string `json:"pointer"`   
}

type Dummy2 struct {
	Name    string  `json:"name"`      //编码为json时，json的这个key为小写的“page”
	Number  int64   `json:"-"`         //编码为json时，忽略自定义的结构体中的这个字段
	Pointer *string `json:"pointer"`   //自定义struct中如果没定义这个field，则编码为json时，取这个field类型的零值（josn中有这个key，但是value为这个key的零值，注意是零值，而不是空“”）
}

type Dummy3 struct {
	Name    string  `json:"name,omitempty"` // 编码为json时，如果自定义的struct中没有这个字段，或者这个字段的值为类型的零值，则编码为json时选择忽略（json中没这个字段）
	Number  int64   `json:"number,omitempty"`
	Pointer *string `json:"pointer,omitempty"`
}

func main() {
	// 基本数据类型到 JSON 字符串的编码过程。这里是一些原子值的例子。
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// 这里是一些切片和 map 编码成 JSON 数组和对象的例子。
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// JSON 包可以自动的编码你的自定义类型。编码仅输出可导出的字段，并且默认使用他们的名字作为 JSON 数据的键。
	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// 你可以给结构字段声明标签来自定义编码的 JSON 数据键名称。在上面 `Response2` 的定义可以作为这个标签的一个例子。
	res2D := Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	//-------------------------------------------------------------------
	// 解码
	// 现在来看看解码 JSON 数据为 Go 值的过程。这里是一个普通数据结构的解码例子。
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)   //json数据

	// 我们需要提供一个 JSON 包可以存放解码数据的变量。这里的 `map[string]interface{}` 将保存一个 string 为键，值为任意值的map。
	var dat map[string]interface{}

	// 这里就是实际的解码和相关的错误检查。
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	fmt.Println(dat["num"])
	fmt.Printf("%T",dat["num"])

	// 为了使用解码 map 中的值，我们需要将他们进行适当的类型转换。例如这里我们将 `num` 的值转换成 `float64` 类型。
	num := dat["num"].(float64)
	fmt.Println(num)

	// 访问嵌套的值需要一系列的转化。
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// 我们也可以解码 JSON 值到自定义类型。这个功能的好处就
	// 是可以为我们的程序带来额外的类型安全加强，并且消除在
	// 访问数据时的类型断言。
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	//例子2
	byt2 := []byte(`
	{
		"name": "Mr Dummy",
		"number": 4,
		"pointer": "yes"
	}
	`)
	var dummy Dummy
	//将json解码为自定义类型dummy
	err := json.Unmarshal(data, &dummy)
	if err != nil {
			fmt.Println("An error occured: %v", err)
			os.Exit(1)
	}
	// we want to print the field names as well
	fmt.Printf("%+v\n", dummy)

	//-----------------------------------------------------------------
	// 在上面的例子中，我们经常使用 byte 和 string 作为使用
	// 标准输出时数据和 JSON 表示之间的中间值。我们也可以和
	// `os.Stdout` 一样，将 JSON 编码直接输出至 `os.Writer`
	// 流中，或者作为 HTTP 响应体。
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)


}
