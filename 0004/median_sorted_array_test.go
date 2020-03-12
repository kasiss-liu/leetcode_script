// 难点在于复杂度 没有解出来啊
package main

import (
	"testing"
)

func findMedianSortedArrays (nums1 []int, nums2 []int) float64 {

	nums := append(nums1,nums2...)
	total := len(nums)

	for i:= 0; i < total; i ++ {
		for j := i; j < total; j++ {
			if nums[i] > nums[j] {
				nums[i],nums[j] = nums[j],nums[i]
			}
		}
	}
	var median float64 

	nlen := len(nums)
	if nlen % 2 == 1 {
		k := int((nlen-1) / 2)
		median = float64(nums[k])
	} else {
		k := nlen / 2 - 1
		median = float64(nums[k] + nums[k+1]) / 2
	}
	return median
}

func TestFindMedianSortedArrays (t *testing.T) {
	n1 := []int{1,3}
	n2 := []int{2,4,4}

	m := findMedianSortedArrays(n1,n2)
	t.Log(m)
}
