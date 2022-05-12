package main

import "fmt"

type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	int
}

func (s *Simple) Set(n int) {
	s.int = n
}

func (s *Simple) Get() int {
	return s.int
}

func simfunc(s Simpler, n int) {
	s.Set(n)
	fmt.Println(s.Get())
}

func main() {
	s := &Simple{1}
	si := Simpler(s)
	si.Set(2)
	fmt.Println(s.Get())
	simfunc(s, 9)
}
