package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		nums           []int
		n              int
		expectedResult []int
	}{
		{
			nums:           []int{2, 5, 1, 3, 4, 7},
			n:              3,
			expectedResult: []int{2, 3, 5, 4, 1, 7},
		},
		{
			nums:           []int{1, 1, 2, 2},
			n:              2,
			expectedResult: []int{1, 2, 1, 2},
		},
	}

	for _, testCase := range testCases {
		result := shuffle(testCase.nums, testCase.n)
		assert.Equal(t, testCase.expectedResult, result)
	}
}
