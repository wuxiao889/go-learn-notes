// parse.go
package parse

import (
	"fmt"
	"strconv"
	"strings"
)

// A ParseError indicates an error in converting a word into an integer.
type ParseError struct {
	Index int    // The index into the space-separated list of words.
	Word  string // The word that generated the parse error.
	Err   error  // The raw error that precipitated this error, if any.
}

//Error和String方法同时实现时，fmt优先调用Error

// String returns a human-readable error message.
func (e *ParseError) Error() string {
	return fmt.Sprintf("Error() pkg parse: error parsing %q as int", e.Word)
}

func (e *ParseError) String() string {
	return fmt.Sprintf("Stirng() pkg parse: error parsing %q as int", e.Word)
}

// Parse parses the space-separated words in in put as integers.
func Parse(input string) (numbers []int, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			//类型断言
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
		}
	}()

	fields := strings.Fields(input)
	numbers = fields2numbers(fields)
	return
}

func fields2numbers(fields []string) (numbers []int) {
	if len(fields) == 0 {
		panic("no words to parse")
	}
	for idx, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			//panic可以接受任意类型的参数
			panic(&ParseError{idx, field, err})
		}
		numbers = append(numbers, num)
	}
	return
}
