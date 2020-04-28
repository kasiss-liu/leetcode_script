package main

import (
	"testing"
)

func lengthOfLIS(nums []int) int {
	maxLen := 0
	ll := len(nums)
	deep := make([]int, ll)

	for i := 0; i < ll; i++ {
		maxval := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if maxval < deep[j] {
					maxval = deep[j]
				}
			}
		}
		deep[i] = maxval + 1
		if maxLen < deep[i] {
			maxLen = deep[i]
		}
	}

	return maxLen
}

func TestLengthOfLIS(t *testing.T) {
	testCase := []int{1, 3, 6, 7, 9, 4, 10, 5, 6}

	ll := lengthOfLIS(testCase)

	t.Log(ll)

}
