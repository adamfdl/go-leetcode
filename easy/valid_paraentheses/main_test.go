package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		input          string
		expectedResult bool
	}{
		{
			input:          "()",
			expectedResult: true,
		},
		{
			input:          "()[]{}",
			expectedResult: true,
		},
		{
			input:          "(]",
			expectedResult: false,
		},
		{
			input:          "([)]",
			expectedResult: false,
		},
		{
			input:          "{[]}",
			expectedResult: true,
		},
		{
			input:          "]",
			expectedResult: false,
		},
		{
			input:          ")(){}",
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		result := isValid(testCase.input)
		assert.Equal(t, result, testCase.expectedResult)
	}
}
