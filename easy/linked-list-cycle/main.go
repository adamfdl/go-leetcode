package main

// TAG: Linked List

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	headNode := &ListNode{Val: 1, Next: nil}
	node2 := &ListNode{Val: 2, Next: nil}
	node3 := &ListNode{Val: 3, Next: nil}
	node4 := &ListNode{Val: 4, Next: headNode}

	appendNode(headNode, node2)
	appendNode(headNode, node3)
	appendNode(headNode, node4)

	if hasCycle(headNode) {
		fmt.Println("Linked list is cyclic")
		return
	}

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

func hasCycle(head *ListNode) bool {
	nodes := make(map[*ListNode]struct{})

	currentNode := head
	for currentNode != nil {
		if _, ok := nodes[currentNode]; ok {
			return true
		} else {
			nodes[currentNode] = struct{}{}
		}
		currentNode = currentNode.Next
	}

	return false
}
