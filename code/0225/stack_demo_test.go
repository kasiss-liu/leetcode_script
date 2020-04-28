package main

import "sync"

type MyStack struct {
	data   []int
	length int
	lock   sync.Mutex
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{data: make([]int, 0, 10), length: 0}
}

/** Push element x onto stack. */
func (ms *MyStack) Push(x int) {
	ms.lock.Lock()
	ms.data = append(ms.data, x)
	ms.length++
	ms.lock.Unlock()
}

/** Removes the element on top of the stack and returns that element. */
func (ms *MyStack) Pop() int {
	ms.lock.Lock()
	ele := ms.Top()
	ms.data = ms.data[:ms.length-1]
	ms.length--
	ms.lock.Unlock()
	return ele
}

/** Get the top element. */
func (ms *MyStack) Top() int {
	return ms.data[ms.length-1]
}

/** Returns whether the stack is empty. */
func (ms *MyStack) Empty() bool {
	if ms.length == 0 {
		return true
	}
	return false
}
