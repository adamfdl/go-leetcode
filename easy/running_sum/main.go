package main

import "fmt"

// TAG: Array

func main() {
	result := runningSum([]int{1, 2, 3})
	fmt.Println(result)
}

func runningSum(nums []int) []int {
	runningSum := []int{}
	for i := 0; i < len(nums); i++ {
		var tempTotal int
		for j := 0; j <= i; j++ {
			tempTotal = tempTotal + nums[j]
		}
		runningSum = append(runningSum, tempTotal)
	}
	return runningSum
}
