package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 脑回路单一的思考方式
func middleNode1(head *ListNode) *ListNode {
	l := 0
	chead := head
	for head != nil {
		l++
		head = head.Next
	}

	if l == 0 {
		return nil
	}

	k := 0
	mid := l / 2
	head = chead
	for head != nil {
		if k == mid {
			return head
		}
		k++
		head = head.Next
	}

	return nil
}

// 奇技淫巧
func middleNode(head *ListNode) *ListNode {
	fastNode := head
	middleNode := head
	for fastNode != nil {
		fastNode = fastNode.Next
		if fastNode != nil {
			fastNode = fastNode.Next
			middleNode = middleNode.Next
		}
	}
	return middleNode
}
