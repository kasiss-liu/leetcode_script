package main

import (
	"fmt"
	"testing"
)

// 删除排序数组中的重复项
// 给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
//
func removeDuplicates(nums []int) int {
	l := len(nums)
	n := 1
	for i := 1; i < l; i++ {
		if nums[i] != nums[i-1] {
			nums[n] = nums[i]
			n++
		}
	}
	fmt.Println(nums)
	return n
}

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	exp := 5

	res := removeDuplicates(nums)

	if exp != res {
		t.Error(exp, "not equal", res)
	}
}
