package main

import (
	"testing"
)

func reverse(x int) int {
	rint := 0
	minInt32 := -(1 << 31)
	maxInt32 := (1 << 31) - 1
	for x != 0 {
		rint = rint*10 + x%10
		if rint > maxInt32 || rint < minInt32 {
			return 0
		}
		x /= 10
	}
	return rint
}

func TestIntReverse(t *testing.T) {
	i := -123

	t.Log(reverse(i))
}
