package main

import "testing"

//MinStack struct
type MinStack struct {
	stack    []int
	minStack []int
}

//Constructor return a new MinStack
func Constructor() MinStack {
	return MinStack{make([]int, 0, 10), make([]int, 0, 10)}
}

//Push push stack
func (ms *MinStack) Push(x int) {
	ms.stack = append(ms.stack, x)

	if len(ms.minStack) > 0 {
		min := ms.minStack[len(ms.minStack)-1]
		if min > x {
			ms.minStack = append(ms.minStack, x)
		} else {
			ms.minStack = append(ms.minStack, min)
		}
	} else {
		ms.minStack = append(ms.minStack, x)
	}

}

//Pop pop stack
func (ms *MinStack) Pop() {
	ms.minStack = ms.minStack[0 : len(ms.minStack)-1]
	ms.stack = ms.stack[0 : len(ms.stack)-1]
}

//Top get top of stack
func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1]
}

//GetMin get min value
func (ms *MinStack) GetMin() int {
	return ms.minStack[len(ms.minStack)-1]
}

//TestMinStack check verify
func TestMinStack(t *testing.T) {
	minStack := Constructor()

	minStack.Push(3)
	minStack.Push(-1)
	minStack.Push(1)
	minStack.Push(-5)

	if minStack.Top() != -5 {
		t.Error("Top() get wrong Value", minStack.Top())
	}

	if minStack.GetMin() != -5 {
		t.Error("GetMin() get wrong Value", minStack.GetMin())
	}

	minStack.Pop()

	if minStack.GetMin() != -1 {
		t.Error("GetMin() get wrong Value")
	}

}
