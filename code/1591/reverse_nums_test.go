package main

import (
	"testing"
)

func reversePairs(nums []int) int {
	return merge(&nums, 0, len(nums)-1)
}

func merge(nums *[]int, left, right int) int {
	if left >= right {
		return 0
	}

	mid := left + (right-left)/2

	cnt := merge(nums, left, mid) + merge(nums, mid+1, right)

	i, j := left, mid+1
	tmp := []int{}

	for i <= mid && j <= right {
		if (*nums)[i] <= (*nums)[j] {
			tmp = append(tmp, (*nums)[i])
			cnt += j - (mid + 1)
			i++
		} else {
			tmp = append(tmp, (*nums)[j])
			j++
		}
	}
	for ; i <= mid; i++ {
		tmp = append(tmp, (*nums)[i])
		cnt += right - (mid + 1) + 1
	}
	for ; j <= right; j++ {
		tmp = append(tmp, (*nums)[j])
	}

	for i := 0; i < len(tmp); i++ {
		(*nums)[left+i] = tmp[i]
	}
	return cnt
}

func TestReversePare(t *testing.T) {
	var nums []int
	var exp, res int

	nums = []int{7, 5, 6, 4}

	res = reversePairs(nums)
	if exp != res {
		t.Error(exp, "not equal", res)
	}
}
