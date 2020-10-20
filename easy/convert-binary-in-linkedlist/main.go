package main

// Tag: Linked List

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node := &ListNode{Val: 1, Next: nil}

	appendNode(node, 0)
	appendNode(node, 1)

	result := getDecimalValue(node)
	fmt.Println(result)
}

func appendNode(node *ListNode, n int) {

	for node.Next != nil {
		node = node.Next
	}

	node.Next = &ListNode{Val: n}
}

func getDecimalValue(node *ListNode) int {
	var binaryString string

	for node != nil {
		binaryString += strconv.Itoa(node.Val)
		node = node.Next
	}

	decimalValue, _ := strconv.ParseInt(binaryString, 2, 64)
	return int(decimalValue)
}
