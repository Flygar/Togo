package main

import "fmt"

/*
有两个类型 stockPosition 和 car，它们都有一个 getValue() 方法，我们可以定义一个具有此方法的接口 valuable。接着定义一个使用 valuable 类型作为参数的函数 showValue()，所有实现了 valuable 接口的类型都可以用这个函数。
*/

type stockPosition struct {
	ticketName string "票名"
	price float64 "价格"
	count int "次数"
}

func (s *stockPosition) getValue() float64 {
	return  s.price*float64(s.count)
}


type car struct {
	carName string
	price float64
}

func (c *car)getValue()  float64{
	return c.price
}

type valuable interface {
	getValue() float64
}

//形参v是指针类型，v代表内存地址。内存地址中的值需要用(*v)来表示。showValue接受的参数是一个地址,变量前加&表示地址。
func showValue(v *valuable)  {
	fmt.Printf("value is : %v\n",(*v).getValue())

}

func main() {
	s0 :=&stockPosition{"动物园",80.5,3}
	c0 :=&car{"AE86",10000}
	v0 :=[]valuable{s0,c0}

	//2种写法实现接口
	//v1 :=s0
	//v2 :=valuable(c0)
	//showValue(v1)
	//showValue(v2)

	for k,_ :=range v0{
		fmt.Printf("%v\n",v0[k].getValue())
		showValue(&(v0[k]))
	}


}
