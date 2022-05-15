package main

import "fmt"

type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	int
}

type RSimple struct {
}

func (s *Simple) Set(n int) {
	s.int = n
}

func (s *Simple) Get() int {
	return s.int
}

func (s *RSimple) Get() int {
	return 0
}

func (s *RSimple) Set(n int) {

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
	fi(si)
	fi(s)
	r := new(RSimple)
	fi(r)
}

func fi(s Simpler) {
	switch s.(type) {
	case *Simple:
		{
			fmt.Println("Simple")
		}
	case *RSimple:
		{
			fmt.Println("Rsimple")
		}
	default:
		{
			fmt.Println("unknow")
		}
	}
}
