package main

import (
	"testing"
)

func myAtoi(s string) int {
	var maxInt = 1<<31 - 1
	var minInt = -1 << 31
	var number = 0
	var symbol = 1
	var first bool = true
	for _, b := range s {
		//非空字符跳过
		if first && b == 32 {
			continue
		}
		//第一个有效字符是符号
		if first && (b == '-' || b == '+') {
			if b == '-' {
				symbol = -1
			}
			first = false
			continue
		}
		//有效int
		if b >= '0' && b <= '9' {
			if first {
				first = false
			}
			number = number*10 + int(b) - 48

			if number*symbol > maxInt {
				return maxInt
			}
			if number*symbol < minInt {
				return minInt
			}
			continue
		}
		//非有效int结束匹配
		if first {
			number = 0
		}
		break
	}

	return number * symbol
}

func TestMyAtoi(t *testing.T) {
	var s string
	var exp, res int

	s = "   -42"
	exp = -42
	res = myAtoi(s)
	if exp != res {
		t.Error(exp, "not equal", res)
	}

	s = "-91283472332"
	exp = -2147483648
	res = myAtoi(s)
	if exp != res {
		t.Error(exp, "not equal", res)
	}

	s = "words and 987"
	exp = 0
	res = myAtoi(s)
	if exp != res {
		t.Error(exp, "not equal", res)
	}

	s = "9223372036854775808"
	exp = 2147483647
	res = myAtoi(s)
	if exp != res {
		t.Error(exp, "not equal", res)
	}
}
