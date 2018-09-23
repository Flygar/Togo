package main

/*
switch
匿名函数
断言
*/

import (
	"fmt"
	"runtime"
	"time"
)

func init() {
	//利用switch判断系统
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os { //也可以加os，下一行： os == "darwin"
	case "darwin":
		fmt.Println("Darwin.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
	fmt.Printf("\n")

}

func main() {
	isWeekend()

	//匿名函数，类型断言
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		case string:
			fmt.Println("I'm a string")
		default:
			fmt.Println(t, "Not bool or int")
		}
	}
	whatAmI(32.7)
	whatAmI("hello")

	//fallthrough
	fmt.Println("fallthrough:")
	testFallthrough()

	//typecheck
	fmt.Println("断言:")
	typecheck(true)
	typecheck("hihi")

	//test
	fmt.Println("test 练习")
	test()
}

//fallthrough：找到符合条件的case后，继续执行后续分支的代码，包括default分支
func testFallthrough() {
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}

}

//isWeekend 判断是否是周末
func isWeekend() {
	switch weekday := time.Now().Weekday(); { // pay attention with the ";" at the end.
	case weekday == time.Saturday || weekday == time.Sunday:
		fmt.Println("take a break")
	case weekday != time.Saturday && weekday != time.Sunday:
		fmt.Println("you need work")
	default:
		fmt.Println("take an err", weekday)
	}
}

//typecheck 类型断言函数
func typecheck(values ...interface{}) {
	for _, value := range values {
		switch value.(type) {
		case int:
			fmt.Println("int", value)
		case float64:
			fmt.Println("float64", value)
		case string:
			fmt.Println("string", value)
		case bool:
			fmt.Println("bool", value)
		default:
			fmt.Println("default", value)
		}
	}
}

//test 练习
func test() {
	for a := 0; a <= 100; a++ {
		switch {
		case a%5 == 0 && a%3 == 0:
			fmt.Println("FizzBuzz")
		case a%5 == 0:
			fmt.Println("Buzz")
		case a%3 == 0:
			fmt.Println("Fizz")
		default:
			fmt.Println(a)
		}
	}
}

//test2 mark
func test2() {
	switch a, b := x[i], y[j]; {
	case a < b:
		t = -1
	case a == b:
		t = 0
	case a > b:
		t = 1
	}
}
