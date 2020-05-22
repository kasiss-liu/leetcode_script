package main

import "testing"

func waysToChange(n int) int {
	ans := 0
	for i := 0; i <= (n / 25); i++ {

		tmp := n - 25*i
		ans = (ans + ((tmp/5+1)+((tmp-10*(tmp/10))/5+1))*(tmp/10+1)/2) % 1000000007
	}
	return ans
}

func TestWaysToChange(t *testing.T) {
	n := 10

	exp := 4

	ret := waysToChange(n)

	if exp != ret {
		t.Error(exp, "not equal", ret)
	}
}
