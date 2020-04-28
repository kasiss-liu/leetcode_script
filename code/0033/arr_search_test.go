package main

import (
	"testing"
)

func search(nums []int, target int) int {
	nlen := len(nums)
	if nlen == 0 {
		return -1
	}
	if nlen == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	left, right := 0, nlen-1

	for left <= right {
		mid := left + (right-left)/2
		if target == nums[mid] {
			return mid
		}

		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

	}
	return -1

}

func TestSearch(t *testing.T) {
	arr := []int{3, 1}
	target := 1

	exp := 4
	res := search(arr, target)
	if exp != res {
		t.Error(exp, "not equal", res)
	}
}
