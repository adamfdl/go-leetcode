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
		tokenPair, isCloseBracket := getTokenPairs[token]
		if isCloseBracket {

			if len(stack) > 0 && stack[len(stack)-1] == tokenPair {
				// If the top of the stack is the opening bracket
				stack.pop()
			} else {
				// If it is not, then it is not valid
				return false
			}
		} else {
			stack.push(token)
		}

	}

	// stack should be empty
	return stack.isEmpty()
}
