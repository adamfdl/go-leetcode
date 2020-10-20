package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func (l *LinkedList) print() {
	currentNode := l.Head
	fmt.Printf("Data: %v\n", currentNode.Val)
	for currentNode.Next != nil {
		currentNode = currentNode.Next
		fmt.Printf("Data: %v\n", currentNode.Val)
	}
}

func (l *LinkedList) prepend(n int) {
	second := l.Head

	newNode := &ListNode{
		Val: n,
	}
	l.Head = newNode
	l.Head.Next = second
}

func (l *LinkedList) appendOrdered(n int) {

	newNode := &ListNode{
		Val: n,
	}

	if l.Head == nil {
		l.Head = newNode
	} else {
		currentNode := l.Head
		if currentNode.Val >= newNode.Val {
			newNode.Next = currentNode
			l.Head = currentNode
			return
		}

		for currentNode.Next != nil && currentNode.Next.Val >= newNode.Val {
			temp := currentNode.Next

			currentNode.Next = newNode
			newNode.Next = temp
			break
		}

		currentNode.Next = newNode
	}
}
