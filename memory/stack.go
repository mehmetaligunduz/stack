package memory

import "fmt"

type Stack struct {
	stack []string
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(item string) {
	if item == "" {
		return
	}
	s.stack = append(s.stack, item)
}

func (s *Stack) Pop() string {
	var index = len(s.stack)
	top := s.stack[index-1]
	s.stack = append(s.stack[:index-1], s.stack[index:]...)
	return top
}

func (s *Stack) ToString() {
	for i := len(s.stack) - 1; i > -1; i-- {
		fmt.Println(s.stack[i])
	}
}
