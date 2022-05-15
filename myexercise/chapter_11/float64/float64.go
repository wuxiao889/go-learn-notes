package float64

import (
	"fmt"
	"math/rand"
	"time"
)

type Float64Array []float64

func (a Float64Array) Len() int           { return len(a) }
func (a Float64Array) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Float64Array) Less(i, j int) bool { return a[i] < a[j] }

func NewFloat64Array() Float64Array {
	return make([]float64, 25)
}

func (a Float64Array) List() string {
	var s string
	for i := range a {
		s += fmt.Sprintf("%v , %1.1f\n", i, a[i])
	}
	return s
}

func (a *Float64Array) String() string {
	return a.List()
}

func (a Float64Array) Fill() {
	rand.Seed(int64(time.Now().Second()))
	for i := range a {
		a[i] = rand.Float64()
	}
}
