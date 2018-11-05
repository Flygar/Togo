package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(ReverseInt(0123))
}

func ReverseInt(i int) int{
	if i> math.MaxInt32 || i < math.MinInt32 {
		return 0
	}
	var resp int
	for i!=0{
		resp = resp*10 + i%10
		i/=10
	}
	return resp

}
