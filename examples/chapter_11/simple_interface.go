package main

import "fmt"

/*
定义一个接口 Simpler，它有一个 Get() 方法和一个 Set()，Get()返回一个整型值，Set() 有一个整型参数。创建一个结构体类型 Simple 实现这个接口
*/

type Simpler interface {
	Get() int
	Set(i int)
}

type Simple struct {
	i int
}

//修改为return
func (s *Simple) Set(a int)  {
	s.i = a
}

//返回一个整形值
func (s *Simple) Get() int {
	return s.i
}

func  fi(s *Simpler)  int {
	(*s).Set(100)
	return (*s).Get()
}
func main() {
	s0 := &Simple{5}

	//var Sr Simpler = &Simple{4}
	//var Sr Simpler = s0
	Sr :=Simpler(s0)
	//Sr := s0

	Sr.Set(4)
	fmt.Printf("%v\n",Sr.Get())

	Sr2 :=Simpler(s0)
	fmt.Printf("%v",fi(&Sr2))
}
