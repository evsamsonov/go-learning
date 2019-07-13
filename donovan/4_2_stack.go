package main

import "fmt"

// Стэк на основе среза
type Stack struct {
	slice []int
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(v int) *Stack {
	s.slice = append(s.slice, v)
	return s
}

func (s *Stack) Pop() int {
	result := s.slice[len(s.slice) - 1]
	s.slice = s.slice[:len(s.slice) - 1]
	return result
}

func main() {
	stack := NewStack()

	for i := 0; i < 10; i++ {
		stack.Push(i)
	}

	fmt.Println(stack)

	for i := 0; i < 5; i++ {
		fmt.Println(stack.Pop())
	}

	fmt.Println(stack)
}


