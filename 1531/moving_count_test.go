package main

import "testing"

func movingCount(m int, n int, k int) int {

	temp := make([][]int, m)
	for a := 0; a < m; a++ {
		temp[a] = make([]int, n)
	}

	count := 1
	temp[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i-1 >= 0 && temp[i-1][j] == 1) || (j-1 >= 0 && temp[i][j-1] == 1) {
				if (add(i) + add(j)) <= k {
					count++
					temp[i][j] = 1
				}
			}
		}
	}
	return count
}

func add(i int) int {
	sum := 0
	for i > 0 {
		sum += i % 10
		i /= 10
	}
	return sum
}

func Test(t *testing.T) {
	var m, n, k, exp, res int

	m = 1
	n = 2
	k = 1
	exp = 2
	res = movingCount(m, n, k)
	if exp != res {
		t.Error(exp, "not equal", res)
	}
}
