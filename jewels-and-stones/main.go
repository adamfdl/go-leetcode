package main

import "fmt"

func main() {
	result := numJewelsInStones("aA", "aAAbbbb")
	fmt.Println(result)
}

func numJewelsInStones(J string, S string) int {
	jewels := []byte(J)
	stones := []byte(S)

	var totalJewels int
	for _, jewel := range jewels {
		for _, stone := range stones {
			if stone == jewel {
				totalJewels++
			}
		}
	}

	return totalJewels
}
