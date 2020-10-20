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
	appendNode(node, 3)
	appendNode(node, 4)
	appendNode(node, 5)

	middleNode := middleNode(node)
	// fmt.Println(result)

	printNodes(middleNode)

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
