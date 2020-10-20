package main

func main() {
	mergeTwoLists(nil, nil)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		newList     = &ListNode{}
		newListHead = newList
	)

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			newList.Next = l1

			l1 = l1.Next
			newList = newList.Next
		} else {
			newList.Next = l2

			l2 = l2.Next
			newList = newList.Next
		}
	}

	if l1 != nil {
		// If there is still data left on l1
		newList.Next = l1
	} else if l2 != nil {
		// If there is still data left on l2
		newList.Next = l2
	}

	return newListHead.Next
}
