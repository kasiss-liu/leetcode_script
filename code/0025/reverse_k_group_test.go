package main

import (
	"fmt"
	"testing"

	. "github.com/kasiss-liu/leetcode_script/code/utils"
)

func TestReverseKGroup(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	k := 2
	list := CreateList(items)
	expect := CreateList([]int{3, 2, 1, 4, 5})

	res := reverseKGroup(list, k)

	extstr := ""
	for expect != nil {
		extstr += fmt.Sprintf("%d ", expect.Val)
		expect = expect.Next
	}
	t.Logf("expect [%s]", extstr)

	resstr := ""
	for res != nil {
		resstr += fmt.Sprintf("%d ", res.Val)
		res = res.Next
	}
	t.Logf("result [%s]", resstr)
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if k <= 1 {
		return head
	}
	var nHead, h, t *ListNode
	p1, p2 := head, head

	for {
		//按k分组
		n := 1
		for p1 != nil && p1.Next != nil && n != k {
			p1 = p1.Next
			n++
		}
		//判断是否读够k个节点
		if n != k {
			if t != nil {
				t.Next = p2
			}
			break
		}
		//记录当前位置cur
		cur := p1
		//p1指针前移
		p1 = p1.Next

		lastT := t
		h, t = reverseList(p2, cur)
		if nHead == nil {
			nHead = h
		}
		if lastT != nil {
			lastT.Next = h
		}
		//p2指针跟随p1前移
		p2 = p1
	}
	return nHead
	// return nil
}

func TestReverseKGroupOpt(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	k := 3
	list := CreateList(items)
	expect := CreateList([]int{3, 2, 1, 4, 5})

	res := reverseKGroupOpt(list, k)

	extstr := ""
	for expect != nil {
		extstr += fmt.Sprintf("%d ", expect.Val)
		expect = expect.Next
	}
	t.Logf("expect [%s]", extstr)

	resstr := ""
	for res != nil {
		resstr += fmt.Sprintf("%d ", res.Val)
		res = res.Next
	}
	t.Logf("result [%s]", resstr)
}

func reverseKGroupOpt(head *ListNode, k int) *ListNode {
	preset := &ListNode{Next: head}
	prev := preset
	for head != nil {
		tail := prev
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return preset.Next
			}
		}
		next := tail.Next
		head, tail = reverseList(head, tail)
		prev.Next = head
		tail.Next = next
		prev = tail
		head = tail.Next
	}
	return preset.Next
}

func TestReversList(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	list := CreateList(items)
	p := list
	n := 1
	for p != nil && n < 3 {
		p = p.Next
		n++
	}
	t.Log("p value ", p.Val)
	start, end := reverseList(list, p)
	t.Logf("expected start 3 : res %d", start.Val)
	t.Logf("expected   end 1 : res %d", end.Val)
}

//将链表指定部分翻转 返回翻转后头尾指针
func reverseList(head, s *ListNode) (start, end *ListNode) {
	end = head
	for head != nil && start != s {
		start, head, head.Next = head, head.Next, start
	}
	return
}
