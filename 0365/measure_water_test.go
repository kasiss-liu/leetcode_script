package main

import "testing"

// 取最大公约数 判断是否整除

func canMeasureWater(x int, y int, z int) bool {
	if x+y < z {
		return false
	}
	if x == 0 || y == 0 {
		return z == 0 || x+y == z
	}
	return z%gcd(x, y) == 0
}

func gcd(i, j int) int {
	if i > j {
		i, j = j, i
	}
	if i == 1 {
		return 1
	}
	y := j % i
	if y == 0 {
		return i
	}
	return gcd(i, y)
}

func TestMeasure(t *testing.T) {
	var exp, res bool

	exp = true
	res = canMeasureWater(3, 4, 5)
	if exp != res {
		t.Error("not equal")
	}
}
