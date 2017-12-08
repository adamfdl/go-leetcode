package main

import (
	"fmt"
)

func main() {
	result := twoSums([]int{3, 2, 3}, 6)
	fmt.Printf("%d", result)
}

func twoSums(nums []int, target int) []int {
	var sum int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {

			} else {
				sum = nums[i] + nums[j]
				if sum == target {
					return []int{i, j}
				}
			}
		}
	}
	return []int{}
}
