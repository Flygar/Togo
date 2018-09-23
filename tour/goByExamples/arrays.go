package main

import "fmt"

func main() {
	//声明数组：通过new 来创建
	var arr1 = new([5]int) //类型是 *[5]int
	fmt.Println(arr1)      //&[0 0 0 0 0]

	//声明数组：定义空 var identifier [len]type
	var arr2 [5]int //类型是 [5]int
	arr2[3] = 333
	fmt.Println(arr2) //[0 0 0 333 0]

	//声明数组并初始化
	var sl = [5]int{1, 2, 3, 4, 5}
	fmt.Println(sl)
	//指定key值的value,string类型的零值为空字符串""
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	fmt.Println(arrKeyValue)

	// 定义数组下标6值为6，下标1值为1
	r := [...]int{6: 6, 1: 1}
	fmt.Println(r)

	//声明数组：并初始化
	a := [...]string{"a", "b", "c", "d"}
	fmt.Println(a)

	//---------------------------------------------------------

	//数组的拷贝
	newArr1 :=[5]string{"H","E","L","L","O"}
	fmt.Println(newArr1)

	//值传递(副本)，对newArr2的修改不会影响newArr1,newArr1的修改也不会影响到newArr2;
	newArr2 :=newArr1
	newArr2[1]="world"
	fmt.Println(newArr2)
	fmt.Println(newArr1)

	//引用传递，newArr3指向newArr1地址上存的值。
	fmt.Println("-----------")
	fmt.Println(newArr1)
	newArr3 := &newArr1
	newArr3[0]="HHHHHHHHH"
	fmt.Println(newArr1)
}
