package main

import "testing"

func canJump(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			isCan := false
			for j := i - 1; j >= 0; j-- {
				if nums[j] > i-j {
					isCan = true
					break
				}
			}
			if !isCan {
				return false
			}
		}
	}
	return true
}

func TestCanJump(t *testing.T) {
	var jumpNums = []int{3, 2, 1, 0, 4}
	var exp, res bool

	exp = false
	res = canJump(jumpNums)
	if exp != res {
		t.Error(exp, "not equal", res)
	}
}
