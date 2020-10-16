package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		nums           []int
		expectedResult []int
	}{
		{
			nums:           []int{6, 5, 4, 8},
			expectedResult: []int{2, 1, 0, 3},
		},
		{
			nums:           []int{7, 7, 7, 7},
			expectedResult: []int{0, 0, 0, 0},
		},
	}

	for _, testCase := range testCases {
		result := smallerNumbersThanCurrent(testCase.nums)
		assert.Equal(t, testCase.expectedResult, result)
	}
}
