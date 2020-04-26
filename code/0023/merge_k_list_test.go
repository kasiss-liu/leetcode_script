package main

import "testing"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//分治思想 简化题目为2个ListNode 合并
func mergeKLists(lists []*ListNode) *ListNode {
	ln := len(lists)
	if ln < 1 {
		return nil
	}
	if ln == 1 {
		return lists[0]
	}

	mid := ln / 2

	l1 := mergeKLists(lists[:mid])
	l2 := mergeKLists(lists[mid:])

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

	for l1 != nil {
		tmp := &ListNode{Val: l1.Val}
		tail.Next = tmp
		tail = tail.Next
		l1 = l1.Next
	}
	for l2 != nil {
		tmp := &ListNode{Val: l2.Val}
		tail.Next = tmp
		tail = tail.Next
		l2 = l2.Next
	}

	return head.Next
}

func createList(a []int) *ListNode {
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
	return head
}

func TestMerge(t *testing.T) {
	a := createList([]int{1, 3, 5})
	b := createList([]int{2, 3, 4})
	// a := createList([]int{})
	// b := createList([]int{})

	lists := []*ListNode{a, b}
	// lists := []*ListNode{}

	res := mergeKLists(lists)
	for res != nil {
		println(res.Val)
		res = res.Next
	}
	t.Error("testing")
}
