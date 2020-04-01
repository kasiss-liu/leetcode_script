package main

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var reverse *ListNode
	for head != nil {
		head, reverse, head.Next = head.Next, head, reverse
	}
	return reverse
}

func TestReverseList(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	head := &ListNode{}
	chain := head
	for k, i := range a {
		if k == 0 {
			chain.Val = i
			continue
		}
		t := &ListNode{Val: i}
		chain.Next = t
		chain = t
	}

	res := reverseList(head)
	chainRes := reverseList(res)

	b := make([]int, 0, len(a))
	for chainRes != nil {
		b = append(b, chainRes.Val)
		chainRes = chainRes.Next
	}

	for k, v := range a {
		if b[k] != v {
			t.Error("not reverse")
		}
	}

}
