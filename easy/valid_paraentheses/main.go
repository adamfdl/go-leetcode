package main

func main() {
	isValid(")(){}")
}

var getTokenPairs = map[string]string{
	"}": "{",
	"]": "[",
	")": "(",
}

func isValid(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return false
	}

	var stack stack

	for i := 0; i < len(s); i++ {
		token := string(s[i])
		openBracket, isCloseBracket := getTokenPairs[token]
		if isCloseBracket {
			if len(stack) > 0 && stack[len(stack)-1] == openBracket {
				stack.pop()
			} else {
				stack.push(token)
			}
		} else {
			stack.push(token)
		}

	}

	return stack.isEmpty()
}

type stack []string

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(str string) {
	*s = append(*s, str)
}

func (s *stack) pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	}

	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}
