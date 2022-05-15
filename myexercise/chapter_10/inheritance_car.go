package main

import "fmt"

type Car struct {
	Engine
	wheelCount int
}

type Engine interface {
	Start()
	Stop()
}

func (this Car) numberOfWheels() int {
	return this.wheelCount
}

type Mercedes struct {
	Car
}

func main() {
	m := new(Mercedes)
	m.wheelCount = 5
	m.Start()
	fmt.Println(m.numberOfWheels())
	m.sayHiToMerkel()
	m.Stop()
	arr := []int{}
	arr = append(arr, 1)
	fmt.Println(arr)
}

func (_ Mercedes) sayHiToMerkel() {
	fmt.Println("hello")
}

func (c *Car) Start() {
	fmt.Println("Car is started")
}

func (c *Car) Stop() {
	fmt.Println("Car is stopped")
}
