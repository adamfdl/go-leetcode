package main

// TAG: Stack

import (
	"bytes"
	"fmt"
)

func main() {

	result := removeDuplicates("abbaca")
	fmt.Println(result)
}

func removeDuplicates(S string) string {
	stack := stack{}

	for i := 0; i < len(S); i++ {
		token := string(S[i])
		if len(stack.items) > 0 && stack.top() == token {
			stack.pop()
		} else {
			stack.push(token)
		}
	}

	duplicateRemoved := bytes.Buffer{}
	for _, val := range stack.items {
		duplicateRemoved.WriteString(val)
	}

	return duplicateRemoved.String()
}

type stack struct {
	items []string
}

func (s *stack) top() string {
	return s.items[len(s.items)-1]
}

func (s *stack) push(item string) {
	s.items = append(s.items, item)
}

func (s *stack) pop() string {
	toRemoveIndex := len(s.items) - 1
	toRemove := s.items[toRemoveIndex]
	s.items = s.items[:toRemoveIndex]
	return toRemove
}
