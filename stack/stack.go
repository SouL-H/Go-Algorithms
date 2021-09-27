package main

import "fmt"

type Stack struct {
	cap   uint64
	depth uint64
	list  []int
}

func inStack(cap uint64) *Stack {
	return &Stack{
		cap:   cap,
		depth: 0,
		list:  make([]int, cap),
	}
}

func (s *Stack) Push(value int) {

	if s.depth > s.cap {
		panic("Kapasite büyük.")
	}
	s.list[s.depth] = value
	s.depth++
}
func (s *Stack) Pop() int {
	if s.depth > 0 {
		s.depth--
		return s.list[s.depth]
	}
	return -1
}

func main() {
	s := inStack(5)
	s.Push(4)
	s.Push(7)
	fmt.Println(s.Pop())
	println(s.Pop())
	println(s.Pop())

}
