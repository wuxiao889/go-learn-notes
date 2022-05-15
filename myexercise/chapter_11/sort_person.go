package main

import (
	"fmt"
	"sort"
)

type Person struct {
	firstname string
	lastname  string
}

type Persons []Person

func (a Persons) Len() int      { return len(a) }
func (a Persons) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Persons) Less(i, j int) bool {
	if a[i].firstname == a[j].firstname {
		return a[i].lastname < a[j].lastname
	}
	return a[i].firstname < a[j].firstname
}

func (a *Persons) String() string {
	return fmt.Sprintf("1")
}

func main() {
	p1 := Person{"peter", "cho"}
	p2 := Person{"jonh", "liu"}
	p3 := Person{"peter", "jiao"}
	ps := Persons{p1, p2, p3}
	fmt.Println(ps)
	sort.Sort(ps)
	fmt.Println(&ps)
}
