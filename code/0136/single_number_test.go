package main

import "testing"

func singleNumber(nums []int) int {
	n := 0
	nLen := len(nums)
	for i := 0; i < nLen; i++ {
		n ^= nums[i]
	}
	return n
}

func TestSingleNumber(t *testing.T) {
	a := []int{1, 2, 3, 3, 2}
	exp := 1

	ret := singleNumber(a)
	if exp != ret {
		t.Error(exp, "not equal", ret)
	}
}
