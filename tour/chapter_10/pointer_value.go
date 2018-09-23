package main

import "fmt"

type st1 struct {
	age int
}

func (p *st1) changeAge()  {
	p.age=50
}

func (p st1) write() string {
	return fmt.Sprint(p.age)
}

func main() {

	st0 :=st1{100}
	st0.changeAge()
	fmt.Println(st0.write())


	st1 := &st1{100}
	st1.changeAge()
	fmt.Println(st1.write())

}
