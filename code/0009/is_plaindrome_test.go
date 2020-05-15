package main

import "testing"

func isPlaindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	y := 0
	n := x
	for x != 0 {
		y = y*10 + (x % 10)
		x = x / 10
	}

	for {
		a := y % 10
		b := n % 10

		if a != b {
			return false
		}

		y /= 10
		n /= 10

		if y == 0 || n == 0 {
			break
		}

	}

	return true

}

func TestIsPlaindrome(t *testing.T) {
	var exp, res bool
	var s int

	s = 12345
	exp = false
	res = isPlaindrome(s)
	if exp != res {
		t.Error(s, "not plaindrome")
	}

	s = 12321
	exp = true
	res = isPlaindrome(s)
	if exp != res {
		t.Error(s, "not plaindrome")
	}

	s = -12345
	exp = false
	res = isPlaindrome(s)
	if exp != res {
		t.Error(s, "not plaindrome")
	}

	s = 0
	exp = true
	res = isPlaindrome(s)
	if exp != res {
		t.Error(s, "not plaindrome")
	}

	s = 121121121121121
	exp = true
	res = isPlaindrome(s)
	if exp != res {
		t.Error(s, "not plaindrome")
	}

	s = 1211211211211211
	exp = true
	res = isPlaindrome(s)
	if exp != res {
		t.Error(s, "not plaindrome")
	}

}
