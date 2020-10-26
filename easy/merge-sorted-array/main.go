package main

import (
	"sort"
)

func main() {
	ar1 := []int{1, 2, 3, 0, 0, 0}
	ar2 := []int{2, 5, 6}

	// ar2 = ar1[2:]

	merge(ar1, 3, ar2, 3)

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2...)
	sort.Ints(nums1)
}
