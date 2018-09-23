package main

import (
	"fmt"
		"math"
	"strconv"
)

func main() {
	//fmt.Printf("%b\n",1)

	a ,b:= "1010","1011"
	fmt.Println(addBinary(a,b))

	
}

func addBinary(a string, b string) string {
	//a2,_:=strconv.Atoi(a)
	//a2b:=fmt.Sprintf("%b",a2)
	//b2,_:=strconv.Atoi(b)
	//b2b:=fmt.Sprintf("%b",b2)

	//fmt.Printf("%T %v | %T %v",a2,a2,b2,b2)
	a2t:=0
	for k,v :=range a{
		vi,_:=strconv.Atoi(string(v))
		a2t+=vi*int(math.Pow(2,float64(len(a)-k-1)))
		//fmt.Println("len",len(a)-k)
	}
	fmt.Println(a2t)

	b2t:=0
	for k,v :=range b{
		vi,_:=strconv.Atoi(string(v))
		b2t+=vi*int(math.Pow(2,float64(len(b)-k-1)))
	}
	fmt.Println(b2t)

	c:=a2t+b2t

	//fmt.Println(c)



	return fmt.Sprintf("%b",c)

}