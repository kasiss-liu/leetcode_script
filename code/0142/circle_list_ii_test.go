package main

import (
	"testing"

	. "github.com/kasiss-liu/leetcode_script/code/utils"
)

func TestCircleNode(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}

	list := CreateCircleList(items, 1)
	// m := make(map[int]struct{})
	// a := list
	// for a != nil {
	// 	t.Log(a.Val)
	// 	if _, ok := m[a.Val]; ok {
	// 		break
	// 	}
	// 	m[a.Val] = struct{}{}
	// 	a = a.Next
	// }
	node := detectCycle(list)

	t.Logf("expected circle val %d", 2)

	t.Logf("result   circle val %d", node.Val)

}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast, slow := head, head
	iscircle := false
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		//相等代表有环
		if slow == fast {
			return getNode(head, slow)
		}
	}
	if !iscircle {
		return nil
	}

	return getNode(head, slow)
}

func getNode(head, slow *ListNode) *ListNode {
	for {
		if head == slow {
			return head
		} else {
			head = head.Next
			slow = slow.Next
		}
	}
}
