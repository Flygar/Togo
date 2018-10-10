package main

import "errors"

type operate func(x, y int) int

func clculate(x, y int, op operate) (int, error) {
	if op ==nil{
		return 0,errors.New("invalid operation")
	}
	return op(x, y), nil
}

func main() {

}
