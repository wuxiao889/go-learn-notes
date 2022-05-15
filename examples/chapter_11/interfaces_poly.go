// interfaces_poly.go
package main

import (
	"fmt"
	"math"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	Shape
	side float32
}

type Circle struct {
	Shape
	// 半径
	redis float32
}

type Shape struct {
}

func (sq *Shape) Area() float32 {
	return 1
}

func (sq *Circle) Area() float32 {
	return sq.redis * sq.redis * math.Pi
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func main() {
	r := Rectangle{5, 3}  // Area() of Rectangle needs a value
	q := &Square{side: 5} // Area() of Square needs a pointer
	shapes := []Shaper{r, q}
	fmt.Println("Looping through shapes for area ...")
	for n := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}

/* Output:
Looping through shapes for area ...
Shape details:  {5 3}
Area of this shape is:  15
Shape details:  &{5}
Area of this shape is:  25
*/
