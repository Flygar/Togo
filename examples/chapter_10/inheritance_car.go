package main

import "fmt"



type Car struct {
	wheelCount int
}

func (m *Car) numberOfWheels() int {
	return m.wheelCount

}

type Mercedes struct {
	Car
	name string
}

func (m *Mercedes) sayHiToMerkel() string  {
	return fmt.Sprint("Hi,"+m.name)
	
}
func main() {
	fmt.Println("---")
	s0 := &Mercedes{Car{4},"lusy"}
	fmt.Println(s0.Car.numberOfWheels())
	fmt.Println(s0.sayHiToMerkel())

}
