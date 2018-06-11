package main

import "fmt"

type List []int

func (l *List) Len() int {
	return len(*l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

//接口内参数返回类型，匿名？
type Appender interface {
	Append(int)
}

//调用接口方法的函数
func CountInfo(a *Appender, start, end int) {
	for i := start; i <= end; i++ {
		(*a).Append(i)
	}

}

type Lener interface {
	Len() int
}

func LongEnough(l *Lener) bool {
	return (*l).Len()*10 > 42

}

func main() {
	//将一个值赋值给一个接口时，编译器会确保所有可能的接口方法都可以在此值上被调用，因此不正确的赋值在编译期就会失败。
	lst := &List{1, 1}
	fmt.Println((lst).Len())
	fmt.Printf("%p,%v",lst,*lst)

	t1:=Appender(lst)

	CountInfo(&t1,2,3)
	fmt.Printf("%v",t1)

	//var t2  Lener = lst
	t2:=Lener(lst)
	LongEnough(&t2)
	fmt.Println(LongEnough(&t2))

	//CountInto(lst, 1, 10)
	//if LongEnough(lst) { // VALID:Identical receiver type
	//	fmt.Printf("- lst is long enough\n")
	//}



}
