package main

import (
	"testing"
)

func convertZArray(s string, numRows int) string {

	length := len(s)
	if length <= numRows || numRows == 1 {
		return s
	}

	ss := make([]byte, length)
	step := numRows + numRows - 2
	idx := 0
	for row := 0; row < numRows; row++ {
		hypo := -1
		if row > 0 && row+1 < numRows {
			hypo = numRows - 1 + numRows - row - 1
		}
		for i := row; i < length; i += step {
			ss[idx] = s[i]
			idx++
			if hypo > 0 && hypo < length {
				ss[idx] = s[hypo]
				idx++
				hypo += step
			}
		}
	}

	return string(ss)
}

func TestConvertZ(t *testing.T) {

	exp := "PINALSIGYAHRPI"
	s := "PAYPALISHIRING"

	ss := convertZArray(s, 4)
	if exp != ss {
		t.Error(exp, " not Equal ", ss)
	}
}
