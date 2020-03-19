package main

import (
	"fmt"
	"math"
	"testing"
)

func reverse(x int) int {
	sybol := 1
	if x < 0 {
		sybol = -1
		x = x * sybol
	}

	ints := make([]int, 0, 10)

	for {
		if x <= 0 {
			break
		}
		c := x % 10
		ints = append(ints, c)
		x = (x - c) / 10
	}
	var rint int
	l := len(ints)
	p := 1
	for i := l - 1; i >= 0; i-- {
		rint = rint + ints[i]*p
		p = p * 10
	}

	if rint > math.MaxInt32 {
		rint = 0
	}

	return rint * sybol
}

func TestIntReverse(t *testing.T) {
	i := -123

	t.Log(reverse(i))
}
