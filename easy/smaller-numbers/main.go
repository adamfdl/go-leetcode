package main

func main() {

}

// TAG: Array

func smallerNumbersThanCurrent(nums []int) []int {
	newArr := []int{}
	for i := 0; i < len(nums); i++ {
		totalSmallerNums := 0
		for j := 0; j < len(nums); j++ {
			if nums[j] == nums[i] {
				continue
			}
			if nums[j] < nums[i] {
				totalSmallerNums++
			}
		}
		newArr = append(newArr, totalSmallerNums)
	}

	return newArr
}
