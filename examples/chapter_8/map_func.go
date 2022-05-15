// map_func.go
package main

import "fmt"

func main() {
	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		5: func() int { return 50 },
	}
	fmt.Println(mf)

	map1 := map[int]int{1: 2, 2: 3, 4: 4}
	fmt.Println(map1)
}

// Output:  map[1:0x10903be0 5:0x10903ba0 2:0x10903bc0]
// Output:  map[1:2 2:3 4:4]
