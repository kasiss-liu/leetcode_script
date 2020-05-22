package main

import (
	"reflect"
	"testing"
)

func plusOne(digits []int) []int {
	ln := len(digits)
	step := 1
	for i := ln - 1; i >= 0; i-- {
		res := digits[i] + step
		if res > 9 {
			step = 1
			digits[i] = res - 10
			continue
		}
		digits[i] = res
		step = 0
		break
	}
	if step == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}

func TestPlusOne(t *testing.T) {
	a := []int{4, 3, 2, 1}

	exp := []int{4, 3, 2, 2}

	ret := plusOne(a)

	if !reflect.DeepEqual(ret, exp) {
		t.Error(exp, "not equal", ret)
	}
}
