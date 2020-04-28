package utils

//ListNode 通用链表结构体
type ListNode struct {
	Val  int
	Next *ListNode
}

//CreateList 使用[]int 创建ListNode
func CreateList(a []int) *ListNode {
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
