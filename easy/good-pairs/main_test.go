package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		nums           []int
		expectedResult int
	}{
		{
			nums:           []int{1, 2, 3, 1, 1, 3},
			expectedResult: 4,
		},
		{
			nums:           []int{1, 1, 1, 1},
			expectedResult: 6,
		},
		{
			nums:           []int{1, 2, 3},
			expectedResult: 0,
		},
	}

	for _, testCase := range testCases {
		result := numIdenticalPairs(testCase.nums)
		assert.Equal(t, testCase.expectedResult, result)
	}
}
