package main

import "testing"

func numberOfSubarrays(nums []int, k int) int {
	ll := len(nums)
	odds := make([]int, ll+1)

	preLen, oddCount := 0, 0

	count := 0
	for i := 0; i <= ll; i++ {
		preLen++
		if i == ll || nums[i]%2 == 1 {
			odds[oddCount] = preLen
			if oddCount >= k {
				count += odds[oddCount-k] * odds[oddCount]
			}
			oddCount++
			preLen = 0
		}
	}
	return count
}

func TestNumberOfSubarrays(t *testing.T) {
	var nums []int
	var k, exp, res int

	nums = []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}
	k = 2
	exp = 16
	res = numberOfSubarrays(nums, k)
	if exp != res {
		t.Error(exp, "not equal", res)
	}
}
