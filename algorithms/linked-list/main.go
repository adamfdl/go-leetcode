package main

import (
	"fmt"
	"strconv"
)

func main() {
	// ll := LinkedList{}
	// ll.appendOrdered(1)
	// ll.appendOrdered(5)
	// ll.appendOrdered(3)
	// ll.appendOrdered(2)
	// ll.print()

	node := &ListNode{Val: 1, Next: nil}

	appendNode(node, 0)
	appendNode(node, 1)

	// for node != nil {
	// 	fmt.Printf("Data: %v\n", node.Val)
	// 	node = node.Next
	// }

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
