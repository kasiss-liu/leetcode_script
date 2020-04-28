package main

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	step1 := head
	step2 := head

	n := 0
	for step1 != nil {
		step1 = step1.Next
		n++
		if step1 != nil {
			n++
			step1 = step1.Next
			step2 = step2.Next
		}
	}

	if n%2 == 1 {
		step2 = step2.Next
	}

	var tmp *ListNode
	for step2 != nil {
		step2, tmp, step2.Next = step2.Next, step2, tmp
	}

	for tmp != nil {
		if tmp.Val != head.Val {
			return false
		}

		tmp = tmp.Next
		head = head.Next
	}

	return true
}

func TestPalindrome(t *testing.T) {
	a := &ListNode{Val: 1}
	b := &ListNode{Val: 2}
	c := &ListNode{Val: 3}
	d := &ListNode{Val: 2}
	e := &ListNode{Val: 1}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e

	t.Error(isPalindrome(a))
}
