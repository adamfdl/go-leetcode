package main

import (
	"fmt"
	"strconv"
)

func main() {
	result := subtractProductAndSum(234)
	fmt.Println(result)
}

// Change bytes to int
// func subtractProductAndSum(n int) int {
// 	strNum := strconv.Itoa(n)

// 	var digits []int
// 	for i := 0; i < len(strNum); i++ {
// 		digit, _ := strconv.Atoi(string(strNum[i]))
// 		digits = append(digits, digit)
// 	}

// 	var product int = 1
// 	var sum int
// 	for _, digit := range digits {
// 		product = product * digit
// 		sum = sum + digit
// 	}

// 	return product - sum
// }

func subtractProductAndSum(n int) int {
	strNum := strconv.Itoa(n)

	var product int = 1
	var sum int

	for i := 0; i < len(strNum); i++ {
		digit, _ := strconv.Atoi(string(strNum[i]))
		product = product * digit
		sum = sum + digit
	}

	return product - sum
}
