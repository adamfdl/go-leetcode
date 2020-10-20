package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	// result := reverse(-1230)
	result := reverse(1534236469)

	fmt.Printf("%d", result)
}

func reverse(x int) int {
	stringOfNum := strconv.Itoa(x)
	stringOfNum = strings.Trim(stringOfNum, "0")

	if isNegative := strings.Contains(stringOfNum, "-"); isNegative {
		stringOfNum = strings.Trim(stringOfNum, "-")
		stringOfNum = stringOfNum + "-"
	}

	var reversed bytes.Buffer
	for i := len(stringOfNum) - 1; i >= 0; i-- {
		char := string(stringOfNum[i])
		reversed.WriteString(char)
	}

	result, _ := strconv.Atoi(reversed.String())

	if result < math.MinInt32 || result > math.MaxInt32 {
		return 0
	}

	return result
}
