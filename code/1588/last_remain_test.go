package main

import (
	"testing"
)

//约瑟夫环。
//参考文章https://www.cnblogs.com/jjscm/p/4463555.html

func lastRemain(n, m int) int {
	p := 0
	for i := 2; i <= n; i++ {
		p = (p + m) % i
	}
	return p
}

func TestLastRemain(t *testing.T) {
	var n, m, exp, res int
	m = 3
	n = 5

	exp = 3
	res = lastRemain(n, m)
	if exp != res {
		t.Error(exp, "not equal", res)
	}
}
