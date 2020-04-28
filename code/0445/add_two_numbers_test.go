package main

import (
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	y1, y2 := make([]int, 0, 5), make([]int, 0, 5)
	ln1, ln2 := 0, 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			y1 = append(y1, l1.Val)
			l1 = l1.Next
			ln1++
		}
		if l2 != nil {
			y2 = append(y2, l2.Val)
			l2 = l2.Next
			ln2++
		}

	}
	maxLn := 0
	if ln1 > ln2 {
		maxLn = ln1
	} else {
		maxLn = ln2
	}

	ints := make([]int, 0, maxLn)
	step := 0
	i := 1
	c := &ListNode{}
	for i <= maxLn {
		sum, c1, c2 := 0, 0, 0
		x1 := ln1 - i
		x2 := ln2 - i
		if x1 >= 0 {
			c1 = y1[x1]
		}
		if x2 >= 0 {
			c2 = y2[x2]
		}
		sum = c1 + c2 + step
		if sum > 9 {
			sum = sum % 10
			step = 1
		} else {
			step = 0
		}
		ints = append(ints, sum)

		i++
		if i > maxLn {
			c.Val = sum
			if step == 1 {
				t := &ListNode{}
				t.Val = step
				t.Next = c
				c = t
			}
			break
		}
		c.Val = sum
		t := &ListNode{}
		t.Next = c
		c = t
	}

	return c
}

func transList(nums1, nums2 []int) *ListNode {
	ln1 := len(nums1)
	ln2 := len(nums2)
	maxLn := 0
	if ln1 > ln2 {
		maxLn = ln1
	} else {
		maxLn = ln2
	}

	ints := make([]int, 0, maxLn)
	step := 0
	i := 0
	for i < maxLn {
		sum, c1, c2 := 0, 0, 0
		x1 := ln1 - i - 1
		x2 := ln2 - i - 1
		if x1 >= 0 {
			c1 = nums1[x1]
		}
		if x2 >= 0 {
			c2 = nums2[x2]
		}
		sum = c1 + c2 + step
		if sum > 9 {
			sum = sum % 10
			step = 1
		} else {
			step = 0
		}
		ints = append(ints, sum)
		i++
	}
	if step > 0 {
		ints = append(ints, step)
	}

	l := &ListNode{}
	c := l
	for i := len(ints) - 1; i >= 0; i-- {
		c.Val = ints[i]
		if i > 0 {
			t := &ListNode{}
			c.Next = t
			c = t
		}
	}
	return l
}

func TestSum(t *testing.T) {
	l1 := transList([]int{5}, []int{})

	l2 := transList([]int{5}, []int{})

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
	t.Error(ints)

}

func TestTransList(t *testing.T) {
	res := transList([]int{7, 2, 4, 3}, []int{5, 6, 4})
	node := res
	ints := make([]int, 0, 5)
	for {
		if node == nil {
			break
		}
		ints = append(ints, node.Val)
		node = node.Next
	}
	t.Error(ints)
}
