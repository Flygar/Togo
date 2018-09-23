package main

import "fmt"

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log  *Log
}
func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

//func (c *Customer) Log() *Log {
//	return c.log
//}

func main() {
	//c := new(Customer)
	//c.Name = "Barak Obama"
	//c.log = new(Log)
	//c.log.msg = "1 - Yes we can!"
	c :=&Customer{"hello",&Log{"1 - Yes we can!"}}
	c.log.Add("2 - After me the world will be a better place!")
	fmt.Println(c.log)

}
