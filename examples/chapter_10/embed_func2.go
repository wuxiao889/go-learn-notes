package main

import (
	"fmt"
)

type Log struct {
	msg string
}

type Customer struct {
	Name string
	Log
}

func main() {
	c := &Customer{"Barak Obama", Log{"1 - Yes we can!"}}
	c.Add("2 - After me the world will be a better place!")
	fmt.Println(c)
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) String() string {
	//值类型log没有实现String()
	return c.Name + "\nLog:\n" + fmt.Sprintln(c.Log)
}

/* Output:
Barak Obama
Log:{1 - Yes we can!
2 - After me the world will be a better place!}
*/

// 接口的实现有关，当值类型实现接口的时候，相当于值类型和该值的指针类型均实现了
// 该接口；相反，当指针类型实现了该接口的时候，只有指针类型实现了接口，值类型是
// 没有实现的。
