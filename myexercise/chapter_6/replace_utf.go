package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	mapping := func(r rune) rune {
		if utf8.RuneLen(r) != 1 {
			r = '?'
		}

		return r
	}
	fmt.Println(strings.Map(mapping, "hello你好"))
	fmt.Println(strings.IndexFunc("hello你好", func(r rune) bool {
		return r == '你'
	}))
	fmt.Println(strings.Map(func(r rune) rune {
		if r > 255 {
			r = '?'
		}
		return r
	}, "jad发aeifi及"))
}
