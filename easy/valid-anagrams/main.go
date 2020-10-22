package main

// TAG: HashTable, Sort
// Other solutions: Create a 26 length array that represent characters in alphabet
// Loop through the first string, increment index that correspond to each character
// Loop through the second string, decrement index that correspond to each character
// Loop through the array, if it is an anagram it should all be zero.

import (
	"fmt"
	"reflect"
)

func main() {
	result := isAnagram("anagram", "nagaram")
	fmt.Println(result)
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hashA := make(map[string]int)
	for i := 0; i < len(s); i++ {
		token := string(s[i])
		if _, ok := hashA[token]; ok {
			hashA[token]++
		} else {
			hashA[token] = 1
		}
	}

	hashB := make(map[string]int)
	for i := 0; i < len(t); i++ {
		token := string(t[i])
		if _, ok := hashB[token]; ok {
			hashB[token]++
		} else {
			hashB[token] = 1
		}
	}

	return reflect.DeepEqual(hashA, hashB)
}
