package main

import (
	"fmt"
	"strconv"
)

type T struct {
	a int
	b float64
	c string
}

func (t *T) String() string {
	return strconv.Itoa(t.a)+" / "+t.c
}

func main() {
	t := &T{7, -2.35, `"abc\tdef"`}
	fmt.Println(t)
	
}
