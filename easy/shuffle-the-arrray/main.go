package main

// TAG: Array

func shuffle(nums []int, n int) []int {
	newArray := []int{}
	for i := 0; i < len(nums)/2; i++ {
		newArray = append(newArray, nums[i])
		newArray = append(newArray, nums[i+len(nums)/2])
	}
	return newArray
}

func main() {

}
