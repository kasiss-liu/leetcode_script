package main

import "testing"

func siftDown(nums []int, index int) {
	nLen := len(nums)
	minIndex := index

	for {
		left, right := 2*index+1, 2*index+2
		if left < nLen && nums[left] < nums[minIndex] {
			minIndex = left
		}
		if right < nLen && nums[right] < nums[minIndex] {
			minIndex = right
		}
		if minIndex == index {
			break
		}
		nums[index], nums[minIndex] = nums[minIndex], nums[index]
		index = minIndex
	}

}

func siftUp(nums []int) {
	nlen := len(nums)

	for parent := nlen / 2; parent >= 0; parent-- {
		heapify(nums, nlen, parent)
	}

}

func heapify(nums []int, nlen int, i int) {
	for {
		left, right := 2*i+1, 2*i+2
		pos := i
		if left < nlen && nums[left] < nums[pos] {
			pos = left
		}
		if right < nlen && nums[right] < nums[pos] {
			pos = right
		}
		if pos == i {
			break
		}

		nums[pos], nums[i] = nums[i], nums[pos]
		i = pos
	}

}

func TestSiftDown(t *testing.T) {
	a := []int{9, 2, 4, 5, 6, 8}

	siftDown(a, 0)

	t.Log(a)

}
func TestSiftUp(t *testing.T) {
	a := []int{9, 2, 4, 5, 6, 8}

	siftUp(a)

	t.Log(a)

}
