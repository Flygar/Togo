package main

import (
	"strings"
	"fmt"
)

/*
结构体定义的三种方式
*/

type person struct{ firstName, lastName string }

func main() {
	var p1 person
	p1.firstName="hello"
	p1.lastName="world"
	upPerson(&p1)
	fmt.Println(p1)

	//p2是指针变量，指向person的内存地址,&person
	p2:=new(person)
	p2.firstName="hello"
	//解指针，反向引用(*&person)
	(*p2).lastName="world"
	upPerson(p2)
	fmt.Println(p2)

	//使用下面的方式定义最优
	p3 := &person{"hello","world"}
	upPerson(p3)
	fmt.Println(p3)


}

func upPerson(p *person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}
