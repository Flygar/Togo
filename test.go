package main

import "fmt"

type phone struct {

}

func (* phone) call() string  {
	return fmt.Sprint("ring")
}
type creame struct {

}

func (* creame) takePhoto() string  {
	return "click"

}

type pc struct {
	phone
	creame
}



func main() {
	c:=&pc{}
	fmt.Println(c.call())
	fmt.Println(c.takePhoto())

}