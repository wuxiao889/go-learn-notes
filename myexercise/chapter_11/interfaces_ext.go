package main

import "fmt"

type AreaInterface interface {
	area() float64
}

func area(a AreaInterface) float64 {
	return a.area()
}

type Triangle struct {
	base   float64
	height float64
}

func (t *Triangle) area() float64 {
	return 0.5 * t.base * t.height
}

type PeriInterface interface {
	Perimeter() float64
}

type Square struct {
	length float64
	height float64
}

func (s *Square) Perimeter() float64 {
	return 2 * (s.length + s.height)
}

func main() {
	t := &Triangle{1, 2}
	s := &Square{1, 3}
	fmt.Println(s.Perimeter())
	fmt.Println(t.area())
	ti := AreaInterface(t)
	fmt.Println(area(ti))
}
