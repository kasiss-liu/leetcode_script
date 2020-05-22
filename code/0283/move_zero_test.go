package main

import (
	"reflect"
	"testing"
)

func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[j] == 0 {
			nums = append(nums[:j], append(nums[j+1:], nums[j])...)
			j--
		}
		j++
	}
}

func TestMoveZero(t *testing.T) {
	a := []int{0, 1, 0, 3, 12}
	exp := []int{1, 3, 12, 0, 0}

	moveZeroes(a)
	if !reflect.DeepEqual(exp, a) {
		t.Error(exp, "not equal", a)
	}

}
