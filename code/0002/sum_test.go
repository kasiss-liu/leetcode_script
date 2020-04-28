package main

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	carry := 0
	v1node := l1
	v2node := l2
	var v1, v2 int
	listNode := &ListNode{}
	currentNode := listNode
	for {
		if v1node == nil && v2node == nil && carry == 0 {
			break
		}

		if v1node != nil {
			v1 = v1node.Val
		} else {
			v1 = 0
		}

		if v2node != nil {
			v2 = v2node.Val
		} else {
			v2 = 0
		}
		node := &ListNode{}
		sum := v1 + v2 + carry
		if sum > 9 {
			carry = 1
			node.Val = sum - 10
		} else {
			carry = 0
			node.Val = sum
		}

		currentNode.Next = node
		currentNode = node

		if v1node != nil {
			v1node = v1node.Next
		}
		if v2node != nil {
			v2node = v2node.Next
		}
	}

	return listNode.Next
}

func TestSum(t *testing.T) {
	l1 := &ListNode{Val: 5}
	l1.Next = &ListNode{Val: 5}

	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 5}

	res := addTwoNumbers(l1, l2)
	node := res
	ints := make([]int, 0, 5)
	for {
		if node == nil {
			break
		}
		ints = append(ints, node.Val)
		node = node.Next
	}
	t.Log(ints)

}
