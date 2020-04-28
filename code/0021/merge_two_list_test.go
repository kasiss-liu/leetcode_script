package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head

	for l1 != nil && l2 != nil {
		tmp := &ListNode{}
		if l1 != nil && l2 != nil {
			if l1.Val <= l2.Val {
				tmp.Val = l1.Val
				l1 = l1.Next
			} else {
				tmp.Val = l2.Val
				l2 = l2.Next
			}
			tail.Next = tmp
			tail = tail.Next
			continue
		}
	}

	if l1 != nil {
		tail.Next = l1
	}
	if l2 != nil {
		tail.Next = l2
	}

	return head.Next
}
