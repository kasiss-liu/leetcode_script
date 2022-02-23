package main

import (
	"fmt"
	"testing"
)

func findMedianSortedArraysNew(nums1 []int, nums2 []int) float64 {
	var a, b []int
	if len(nums1) > len(nums2) {
		a = nums1
		b = nums2
	} else {
		a = nums2
		b = nums1
	}

	for _, n := range b {
		start, end := 0, len(a)-1
		for {
			mid := (start + end) / 2
			if n <= a[start] {
				a = append(a[0:start], append([]int{n}, a[start:]...)...)
				break
			}
			if n >= a[end] {
				a = append(a[0:end+1], append([]int{n}, a[end+1:]...)...)
				break
			}
			if n > a[start] && n <= a[mid] {
				end = mid - 1
			}
			if n > a[mid] && n < a[end] {
				start = mid + 1
			}
		}
	}
	fmt.Println(a)
	if len(a)%2 == 1 {
		return float64(a[len(a)/2])
	} else {
		mid := len(a) / 2
		fmt.Println(a[mid], a[mid-1], mid)
		return float64(a[mid]+a[mid-1]) / 2
	}
}

func TestFindMedianSortedArraysNew(t *testing.T) {
	n1 := []int{3}
	n2 := []int{-1, -2}

	// n1 := []int{0, 0, 0, 0, 0}
	// n2 := []int{-1, 0, 0, 0, 0, 0, 1}

	m := findMedianSortedArraysNew(n1, n2)
	t.Log(m)
}
