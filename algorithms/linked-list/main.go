package main

import "fmt"

func main() {
	// ll := LinkedList{}
	// ll.appendOrdered(1)
	// ll.appendOrdered(5)
	// ll.appendOrdered(3)
	// ll.appendOrdered(2)
	// ll.print()

	node := &ListNode{Val: 1, Next: nil}
	appendNode(node, 2)
	appendNode(node, 2)
	appendNode(node, 3)
	appendNode(node, 4)
	appendNode(node, 5)

	// middleNode := middleNode(node)
	// printNodes(middleNode)

	// fmt.Println("---")
	// newHeadNode := reverseList(node)

	// for newHeadNode != nil {
	// 	fmt.Printf("Data: %v\n", newHeadNode.Val)
	// 	newHeadNode = newHeadNode.Next
	// }

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

func hasCycle(head *ListNode) bool {
	var (
		nodes map[interface{}]struct{}
	)

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

// https://www.youtube.com/watch?v=sYcOK51hl-A&ab_channel=mycodeschool
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

func appendNode(node *ListNode, n int) {

	for node.Next != nil {
		node = node.Next
	}

	node.Next = &ListNode{Val: n}
}

func printNodes(node *ListNode) {
	for node != nil {
		fmt.Printf("Data: %v\n", node.Val)
		node = node.Next
	}
}

func middleNode(head *ListNode) *ListNode {

	var (
		tempHeadNode = head
		length       int
	)

	for head != nil {
		length++
		head = head.Next
	}

	var middleNode int = (length / 2)
	for i := 0; i < middleNode; i++ {
		fmt.Println(tempHeadNode.Val)
		tempHeadNode = tempHeadNode.Next
	}

	return tempHeadNode
}
