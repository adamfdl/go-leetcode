package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{12, 345, 2, 6, 7896}
	result := findNumbers(nums)

	fmt.Println(result)
}

func findNumbers(nums []int) int {
	var total int
	for _, num := range nums {
		strNum := strconv.Itoa(num)
		if len(strNum)%2 == 0 {
			total++
		}
	}
	return total
}
