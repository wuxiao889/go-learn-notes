package main

import (
	"fmt"
	"sort"

	"./float64"
)

func main() {
	fa := float64.NewFloat64Array()
	fa.Fill()
	fmt.Println(&fa)
	sort.Sort(fa)
	fmt.Println(&fa)
}
