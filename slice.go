package main

/*
数组是值类型，可以通过new()来创建
*/
import "fmt"

func main() {
	//区别：arr1 的类型是 *[5]int(指针，指向内存地址)，而 arr2的类型是 [5]int。
	var arr1 = new([5]int)
	var arr2 = [5]int{1, 1, 1, 1, 5}
	//arr1作为指针，指向arr2的内存地址;对arr1或arr2做出的修改都会改变内存中存储的值,即arr1和arr2相互影响。
	//指针(arr1)的用法：指向变量的地址(取址符&)，所以arr2前面得加上&号
	arr1 = &arr2
	arr1[1] = 2
	arr1[2] = 3
	arr2[4] = 666
	//不在arr1前加*号，指向的也是内存地址
	fmt.Println(*arr1, "\n", arr2)
	fmt.Println("-----华丽的分割线------")
	arr1[0] = 50
	arr1[1] = 51
	arr1[2] = 52
	arr1[3] = 53
	fmt.Println("arr1:", arr1)
	//重点!：当把一个数组赋值给另一个时，会做一次数组内存的拷贝操作
	arr2 = *arr1
	//arr2是arr1的拷贝，修改arr2不会对arr1产生影响，因为arr2是arr1的拷贝
	arr2[0] = 100
	fmt.Println("arr2:", arr2)

	fmt.Println("------狗程爽的分割线-------")
	var arr3 = [5]int{1, 2, 3, 4}
	//arr4 是 arr3 的拷贝，对arr4做出的修改不会影响arr3
	arr4 := arr3
	arr4[0] = 100
	fmt.Println(arr3, arr4)

	fmt.Println("------slice和array的区别-------")

	//不推荐，消耗更多的内存，拷贝了数组
	var att2 = [3]int{1, 2, 3} //数组
	fmt.Println("-------数组--------")
	fmt.Println(att2) // [1 2 3]
	ftt2(att2)        // 数组，产生一次拷贝，不会修改原数组。
	fmt.Println(att2) // [1 2 3]

	//推荐用法，切片，引用传递，不会拷贝数组
	var att = []int{1, 2, 3}   //切片

	fmt.Println("------切片------")
	fmt.Println(att) // [1 2 3]
	ftt(att)         // 切片，引用传递，修改原始的数组 att。
	fmt.Println(att) // [111 2 3]


	//fp(&ar) // 指针，指向ar数组，fp 方法会修改原始数组ar

	fmt.Println("-------------")
	for i := 0; i < 3; i++ {
		ftt([]int{i, i * i, i * i * i})
		fp(&[3]int{i, i * i, i * i * i})
	}

	//Sum
	fmt.Println("-------第三条华丽的线-------")
	array := [3]float64{7.0, 8.5, 9.1}
	x := Sum(&array) // Note the explicit address-of operator
	// to pass a pointer to the array
	fmt.Printf("The sum of the array is: %f", x)
}

func ftt(a []int) {  //切片，引用传递，会修改原来数组
	a[0] = 111
	fmt.Println(a)
}

func ftt2(a [3]int) { //数组，产生一次拷贝，不会修改原数组
	a[0] = 222
	fmt.Println(a)
}

func fp(a *[3]int) {
	//&a[1]=888
	fmt.Println(a)
}

//把一个大数组传递给函数会消耗很多内存。有两种方法可以避免这种现象：
//传递数组的指针
//使用数组的切片
func Sum(a *[3]float64) (sum float64) {
	for _, v := range a { // derefencing *a to get back to the array is not necessary!
		sum += v
	}
	return
}
