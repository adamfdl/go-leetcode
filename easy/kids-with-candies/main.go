package main

// TAG: Array

func main() {
	kidsWithCandies([]int{2, 3, 5, 1, 3}, 3)
}

func kidsWithCandies(candies []int, extraCandies int) []bool {

	greatestNumCandies := 0
	for i := 0; i < len(candies); i++ {
		if candies[i] > greatestNumCandies {
			greatestNumCandies = candies[i]
		}
	}

	canHaveGreatest := []bool{}
	for i := 0; i < len(candies); i++ {
		if candies[i] == greatestNumCandies {
			canHaveGreatest = append(canHaveGreatest, true)
		} else if candies[i]+extraCandies >= greatestNumCandies {
			canHaveGreatest = append(canHaveGreatest, true)
		} else {
			canHaveGreatest = append(canHaveGreatest, false)
		}
	}

	return canHaveGreatest
}
