package main

import "testing"

func subarraySum(nums []int, k int) int {
	count := 0
	nLen := len(nums)
	for i := 0; i < nLen; i++ {
		sum := 0
		for j := i; j >= 0; j-- {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}
	return count
}

func TestSumArray(t *testing.T) {
	a := []int{1, 1, 1}
	k := 2

	exp := 2
	ret := subarraySum(a, k)

	if exp != ret {
		t.Error(exp, "not equal", ret)
	}
}
