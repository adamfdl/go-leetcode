package main

import "fmt"

// TAG: Linked List

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	headNode := &ListNode{Val: 1, Next: nil}
	node2 := &ListNode{Val: 2, Next: nil}
	node3 := &ListNode{Val: 3, Next: nil}

	appendNode(headNode, node2)
	appendNode(headNode, node3)

	deleteNode(node2)

	for headNode != nil {
		fmt.Printf("Data: %v\n", headNode.Val)
		headNode = headNode.Next
	}
}

func appendNode(node *ListNode, newNode *ListNode) {
	for node.Next != nil {
		node = node.Next
	}

	node.Next = newNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
