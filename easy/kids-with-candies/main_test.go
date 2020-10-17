package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		candies        []int
		extraCandies   int
		expectedResult []bool
	}{
		{
			candies:        []int{2, 3, 5, 1, 3},
			extraCandies:   3,
			expectedResult: []bool{true, true, true, false, true},
		},
		{
			candies:        []int{4, 2, 1, 1, 2},
			extraCandies:   1,
			expectedResult: []bool{true, false, false, false, false},
		},
		{
			candies:        []int{12, 1, 12},
			extraCandies:   10,
			expectedResult: []bool{true, false, true},
		},
	}

	for _, testCase := range testCases {
		result := kidsWithCandies(testCase.candies, testCase.extraCandies)
		assert.Equal(t, testCase.expectedResult, result)
	}
}
