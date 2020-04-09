package main

import "testing"

//数组升序排序
func sortBubbleArray(nums []int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

func sortQuickArray(nums []int) []int {
	l := len(nums)
	if l <= 1 {
		return nums
	}

	var sort func(left, right int)
	sort = func(l, r int) {
		if r-l < 1 {
			return
		}
		base := nums[l]
		left := l + 1
		right := r

		for {
			if left >= right {
				break
			}

			for right >= l+1 && nums[right] >= base {
				right--
			}
			for left <= r && nums[left] < base {
				left++
			}

			nums[left], nums[right] = nums[right], nums[left]
		}
		nums[l], nums[right] = nums[right], nums[l]
		sort(l, right-1)
		sort(right+1, r)
	}

	sort(0, len(nums)-1)

	return nums
}

func TestSortArray(t *testing.T) {
	nums := make([]int, 500, 500)
	for i := 499; i >= 0; i-- {
		nums[i] = i
	}

	res := sortQuickArray(nums)
	t.Log(res[0], res[499])
}
