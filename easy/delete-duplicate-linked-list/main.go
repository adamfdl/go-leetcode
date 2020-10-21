package main

// Tag: Linked List

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node := &ListNode{Val: 1, Next: nil}
	appendNode(node, 2)
	appendNode(node, 2)
	appendNode(node, 3)
	appendNode(node, 4)
	appendNode(node, 5)

	updatedNode := deleteDuplicates(node)

	for updatedNode != nil {
		fmt.Printf("Data: %v\n", updatedNode.Val)
		updatedNode = updatedNode.Next
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	nodes := make(map[int]struct{})

	newNode := &ListNode{}

	currentNode := head
	for currentNode != nil {
		if _, ok := nodes[currentNode.Val]; !ok {
			nodes[currentNode.Val] = struct{}{}
			appendNode(newNode, currentNode.Val)
		}
		currentNode = currentNode.Next
	}

	return newNode.Next
}

func appendNode(node *ListNode, n int) {

	for node.Next != nil {
		node = node.Next
	}

	node.Next = &ListNode{Val: n}
}
