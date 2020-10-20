package main

import (
	"bytes"
	"fmt"
)

var arrOfMorse = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func main() {
	result := uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"})
	fmt.Println(result)
}

func uniqueMorseRepresentations(words []string) int {
	res := map[string]bool{}
	for _, word := range words {
		var morse bytes.Buffer
		for _, char := range word {
			morse.WriteString(arrOfMorse[char-'a'])
		}
		res[morse.String()] = true
	}
	return len(res)
}
