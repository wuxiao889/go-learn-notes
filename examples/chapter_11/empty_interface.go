package main

import "fmt"

var i = 5
var str = "ABC"

type Person struct {
	name string
	age  int
}

type Any interface{}

func main() {
	var val Any
	val = 5
	f := func(t Any) {
		switch t := val.(type) {
		case int:
			fmt.Printf("Type int %v\n", t)
		case string:
			fmt.Printf("Type string %v\n", t)
		case bool:
			fmt.Printf("Type boolean %v\n", t)
		case *Person:
			fmt.Printf("Type pointer to Person %v\n", *t)
		default:
			fmt.Printf("Unexpected type %T", t)
		}
	}
	f(val)
	val = str
	f(val)
	pers1 := new(Person)
	pers1.name = "Rob Pike"
	pers1.age = 55
	val = pers1
	f(val)
}
