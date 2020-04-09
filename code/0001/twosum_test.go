package main

import "testing"

func twoSum(nums []int, target int) []int {
	res := make([]int, 0, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res = append(res, i, j)
				return res
			}
		}
	}
	return res
}

func TestTwoSum(t *testing.T) {
	nums := make([]int, 0, 6)
	nums = append(nums, -1, -2, -3, -4, -5)
	target := -8

	res := twoSum(nums, target)
	t.Log(res)
}
