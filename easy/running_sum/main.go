package main

import "fmt"

// TAG: Array

func main() {
	result := runningSum([]int{1, 2, 3})
	fmt.Println(result)

	result = runningSumV2([]int{1, 2, 3})
	fmt.Println(result)
}

// From discussion.
// Uses only 1 for loop.
// What it does: Add total of n-1 + n
func runningSumV2(nums []int) []int {
	runningSum := []int{}
	runningSum = append(runningSum, nums[0])

	for i := 1; i < len(nums); i++ {
		tempTotal := runningSum[i-1] + nums[i]
		runningSum = append(runningSum, tempTotal)
	}

	return runningSum
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
