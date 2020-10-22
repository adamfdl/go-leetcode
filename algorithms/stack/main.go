package main

import "fmt"

func main() {
	stack := stack{}
	stack.push(1)
	stack.push(1)
	stack.push(1)
	stack.push(1)
	stack.push(1)

	fmt.Println(stack.items)

	stack.pop()

	fmt.Println(stack.items)
}

type stack struct {
	items []int
}

func (s *stack) push(item int) {
	s.items = append(s.items, item)
}

func (s *stack) pop() int {
	toRemoveIndex := len(s.items) - 1
	toRemove := s.items[toRemoveIndex]
	s.items = s.items[:toRemoveIndex]
	return toRemove
}
