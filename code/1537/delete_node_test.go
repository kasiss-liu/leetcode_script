package main

import (
	"reflect"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {

	chain := head
	var prev *ListNode
	for chain != nil {
		if chain.Val == val {
			if prev == nil {
				head = chain.Next
			} else {
				prev.Next = chain.Next
			}
			break
		}
		prev = chain
		chain = chain.Next

	}
	return head
}

func TestDeleteNode(t *testing.T) {
	var head *ListNode
	var vals = []int{4, 5, 1, 9}
	head = &ListNode{Val: vals[0]}
	chain := head
	for i := 1; i < len(vals); i++ {
		temp := &ListNode{Val: vals[i]}
		chain.Next = temp
		chain = temp
	}

	var exp = []int{4, 1, 9}
	var del = 5
	var res = make([]int, 0, len(exp))
	resHead := deleteNode(head, del)
	for resHead != nil {
		res = append(res, resHead.Val)
		resHead = resHead.Next
	}

	if !reflect.DeepEqual(exp, res) {
		t.Error(exp, "not equal", res)
	}

}
