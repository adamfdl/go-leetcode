package main

func main() {

}

// TAG: Array

func numIdenticalPairs(nums []int) int {
	identicalPairs := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i < j && nums[i] == nums[j] {
				identicalPairs++
			}
		}
	}
	return identicalPairs
}
