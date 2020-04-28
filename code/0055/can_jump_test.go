package main

import "testing"

//如果遇到0 则判断前面的步长能否越过本次节点 如果不能 则永远到达不了终点
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

//根据可到达位置判断
func canJumpMax(nums []int) bool {
	maxPos := 0
	ln := len(nums)
	for i := 0; i < ln; i++ {
		if maxPos < i {
			return false
		}

		maxPos = max(maxPos, i+nums[i])
		if maxPos >= ln-1 {
			return true
		}

	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestCanJump(t *testing.T) {
	var jumpNums = []int{3, 2, 1, 0, 4}
	var exp, res bool

	exp = false
	res = canJumpMax(jumpNums)
	if exp != res {
		t.Error(exp, "not equal", res)
	}
}
