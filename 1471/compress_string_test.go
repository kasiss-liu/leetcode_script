package main

import (
	"strconv"
	"testing"
)

func compressString(s string) string {

	bs := []byte(s)
	ll := len(bs)

	if ll < 1 {
		return s
	}

	nbs := make([]byte, 0, ll)
	num := 1
	lastChar := bs[0]
	for i := 1; i < ll; i++ {
		if bs[i] == lastChar {
			num++
			continue
		}

		nbs = append(nbs, lastChar)
		nbs = append(nbs, []byte(strconv.Itoa(num))...)

		lastChar = bs[i]
		num = 1
	}
	nbs = append(nbs, lastChar)
	nbs = append(nbs, []byte(strconv.Itoa(num))...)

	if len(nbs) < ll {
		return string(nbs)
	}
	return s

}

func TestCompressString(t *testing.T) {
	var s, res, exp string

	s = "aaaaabbbddcceedd"
	exp = "a5b3d2c2e2d2"

	res = compressString(s)
	if exp != res {
		t.Error(exp, " != ", res)
	}

}
