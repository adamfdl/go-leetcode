package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node := &ListNode{Val: 1, Next: nil}
	appendNode(node, 2)
	appendNode(node, 3)
	appendNode(node, 4)
	appendNode(node, 5)

	newHeadNode := reverseList(node)
	for newHeadNode != nil {
		fmt.Printf("Data: %v\n", newHeadNode.Val)
		newHeadNode = newHeadNode.Next
	}
}

func appendNode(node *ListNode, n int) {

	for node.Next != nil {
		node = node.Next
	}

	node.Next = &ListNode{Val: n}
}

func reverseList(head *ListNode) *ListNode {
	var (
		current *ListNode = head
		prev    *ListNode
		next    *ListNode
	)

	for current != nil {
		next = current.Next
		current.Next = prev

		prev = current
		current = next
	}

	// Last `previous` node should be the new head
	head = prev
	return head
}
