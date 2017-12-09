package main

import "fmt"
import "strconv"

func main() {
	result := isPalindrome(10)
	fmt.Println(result)
}

func isPalindrome(num int) bool {
	numString := strconv.Itoa(num)

	right := len(numString) - 1
	for i := 0; i < len(numString); i++ {
		left := i

		if left < right {
			if string(numString[left]) != string(numString[right]) {
				return false
			}
		}

		right--
	}

	return true
}
