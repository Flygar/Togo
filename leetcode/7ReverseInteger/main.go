package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(reverseDigits(-3200))
//a:=reverse(9000)
//fmt.Println(a)
}

func reverseDigits(num int) (ret int) {
	ret = 0
	var err error
	var byteData []byte
	//不在32位数的范围内，就return int类型的零值0
	if num > math.MaxInt32 || num < math.MinInt32 {
		return
	}
	switch {
	case num >= 0:
		str1 := strconv.Itoa(num)
		//str1:=[]byte(string(num))
		for i:=len(str1)-1;i>=0;i--{
			//fmt.Printf("%c\n",str1[i])
			byteData=append(byteData,byte(str1[i]))
		}
		ret,err=strconv.Atoi(string(byteData))
		if err!=nil{
			panic(err)
		}
		//return ret
	case num < 0:
		//str1:=string(num)
		str1:=strconv.Itoa(num)
		fmt.Println(str1)
		//if
		for i:=len(str1)-1;i>=0;i--{
			byteData=append(byteData,byte(str1[i]))
		}
		if string(byteData[len(byteData)-1]) == "-" {
			byteData = byteData[:len(byteData)-1]
			byteData = append([]byte("-"), byteData...)
		}
		ret,err=strconv.Atoi(string(byteData))
		if err!=nil{
			panic(err)
		}
		//fmt.Println(string(byteData))

	}
	return

}

func reverse(x int) int {
	xString := strconv.Itoa(x)
	var byteData []byte
	for i := len(xString) - 1; i >= 0; i-- {
		byteData = append(byteData, byte(xString[i]))
	}
	if string(byteData[len(byteData)-1]) == "-" {
		byteData = byteData[:len(byteData)-1]
		byteData = append([]byte("-"), byteData...)
	}
	n, err := strconv.Atoi(string(byteData))
	if err != nil {
		panic(err)
	}
	if n > math.MaxInt32 || n < math.MinInt32 {
		return 0
	}
	return n
}
