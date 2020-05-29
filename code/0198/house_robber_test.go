package main

import "testing"

func rob(nums []int) int {
	nLen := len(nums)
	if nLen == 0 {
		return 0
	}
	if nLen == 1 {
		return nums[0]
	}

	var f1, f2 int
	f1 = nums[0]
	f2 = max(nums[0], nums[1])

	for i := 2; i < nLen; i++ {
		f1, f2 = f2, max(f2, f1+nums[i])
	}
	return f2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestRobber(t *testing.T) {
	houses := []int{1, 2, 3, 1}
	exp := 4

	ret := rob(houses)
	if exp != ret {
		t.Error(exp, "not equal", ret)
	}

}
