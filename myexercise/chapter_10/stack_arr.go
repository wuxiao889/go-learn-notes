package main

import "fmt"

type stack struct {
	arr []int
}

func (s *stack) push(n int) {
	s.arr = append(s.arr, n)
}

func (s *stack) poll() int {
	if len(s.arr) == 0 {
		return 0
	}
	i := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return i
}

func (s stack) String() string {
	str := ""
	for i := range s.arr {
		str += fmt.Sprintf("%v %v\n", i, s.arr[i])
	}
	return str
}

func main() {
	s := stack{[]int{}}
	s.push(1)
	s.push(2)
	s.push(2)
	s.poll()
	fmt.Println(s)
}
