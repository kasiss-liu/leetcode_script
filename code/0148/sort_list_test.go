package main

import (
	"testing"

	. "leetcode.kasiss.cn/code/utils"
)

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	right := slow.Next
	slow.Next = nil
	return mergeTwoLists(sortList(head), sortList(right))
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
func TestMerge(t *testing.T) {
	a := CreateList([]int{6, 5, 3, 2, 10, 3, 2, 5})
	res := sortList(a)
	t.Error("testing", res)
}
