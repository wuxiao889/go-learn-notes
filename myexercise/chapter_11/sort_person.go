package main

import (
	"fmt"
	"sort"
	"strings"
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

func (a *Persons) String() (s string) {
	for i, v := range *a {
		s += fmt.Sprintf("person %v:", i)
		s += fmt.Sprintf("%s%s", strings.ToUpper(v.firstname[:1]), v.firstname[1:])
		s += fmt.Sprintf("%s%s\n", strings.ToUpper(v.lastname[:1]), v.lastname[1:])
	}
	return
}

func main() {
	p1 := Person{"peter", "cho"}
	p2 := Person{"jonh", "liu"}
	p3 := Person{"peter", "jiao"}
	ps := &Persons{p1, p2, p3}
	fmt.Printf("%v", ps)
	sort.Sort(ps)
	fmt.Printf("after sort:\n%v", ps)
}
