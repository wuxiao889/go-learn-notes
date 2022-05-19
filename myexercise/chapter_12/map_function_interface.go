package main

import "fmt"

func main() {
	s := []any{1, 2, 3, "jfalk", "name"}
	mf := func(v any) any {
		switch t := v.(type) {
		case int:
			return t * 2
		case string:
			return t + t
		}
		return v
	}
	fmt.Println(mapFunc(mf, s))
}

func mapFunc(mf func(a any) any, list []any) []any {
	result := make([]any, len(list))
	for i := range list {
		result[i] = mf(list[i])
	}
	return result
}
