// sort_map.go
// the telephone alphabet:
package main

import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23, "delta": 87,
		"echo": 56, "foxtrot": 12, "golf": 34, "hotel": 16, "indio": 87, "juliet": 65, "kilo": 43, "lima": 98}
)

func main() {
	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}
	//创建一个key的切片然后拷贝key
	keys := make([]string, len(barVal))
	i := 0
	for k := range barVal {
		keys[i] = k
		i++
	}

	//对切片排序
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")

	//使用切片的for-range打印所有键值对
	//此时k(value)是key
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v / ", k, barVal[k])
	}
	println()
}

/* Output:
unsorted:
Key: indio, Value: 87 / Key: echo, Value: 56 / Key: juliet, Value: 65 / Key: charlie, Value: 23 /
Key: hotel, Value: 16 / Key: lima, Value: 98 / Key: bravo, Value: 56 / Key: alpha, Value: 34 /
Key: kilo, Value: 43 / Key: delta, Value: 87 / Key: golf, Value: 34 / Key: foxtrot, Value: 12 /
sorted:
Key: alpha, Value: 34 / Key: bravo, Value: 56 / Key: charlie, Value: 23 / Key: delta, Value: 87 /
Key: echo, Value: 56 / Key: foxtrot, Value: 12 / Key: golf, Value: 34 / Key: hotel, Value: 16 /
Key: indio, Value: 87 / Key: juliet, Value: 65 / Key: kilo, Value: 43 / Key: lima, Value: 98 /
*/
