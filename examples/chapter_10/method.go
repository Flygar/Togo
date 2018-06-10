package main

import "fmt"

/*
结构体类型方法
*/

type twoNum struct {
	a,b int
}
func main() {
	t0 := &twoNum{1,2}
	//调用函数
	fmt.Println(addSum(t0))
	//调用方法0
	fmt.Println(t0.addSum())
	//调用方法1
	fmt.Println(t0.addWith(3))

	//指针方法和值方法都可以在指针或非指针上被调用
	//对于类型 T，并且 t 是这个类型的变量,如果在 *T 上存在方法 Meth()，那么 t.Meth() 会被自动转换为 (&t).Meth()
	t1:=twoNum{3,4}
	(&t1).addSum()
	
}

//函数
func addSum(p *twoNum) int {
	return p.a+p.b
}

//方法0
func (p *twoNum) addSum() (s int) {
	s=p.b+p.a
	return s

}
//方法1
func (p *twoNum) addWith(parm int) (s int){
	return p.addSum()+parm
}
