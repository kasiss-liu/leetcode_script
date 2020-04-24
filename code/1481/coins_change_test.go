package main

import "testing"

func waysToChange(n int) int {
	ans := 0
	for i := 0; i <= (n / 25); i++ {
		tmp := n - 25*i
		ans += ((tmp/5 + 1) + ((tmp-10*(tmp/10))/5 + 1)) * (tmp/10 + 1) / 2
	}
	return ans
}

func TestCoinsChange(t *testing.T) {
	a := 1000000
	res := waysToChange(a)
	t.Error(res)
}
