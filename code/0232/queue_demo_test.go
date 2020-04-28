package main

type MyQueue struct {
	data   []int
	length int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{data: make([]int, 0, 10), length: 0}
}

/** Push element x to the back of queue. */
func (mq *MyQueue) Push(x int) {
	mq.data = append(mq.data, x)
	mq.length++
}

/** Removes the element from in front of queue and returns that element. */
func (mq *MyQueue) Pop() int {
	ele := mq.Peek()
	mq.data = mq.data[1:]
	mq.length--
	return ele
}

/** Get the front element. */
func (mq *MyQueue) Peek() int {
	return mq.data[0]
}

/** Returns whether the queue is empty. */
func (mq *MyQueue) Empty() bool {
	if mq.length == 0 {
		return true
	}
	return false
}
