//暴力破解的本办法
package main

import (
	"testing"
)

type charPos struct {
	lastIndex    int
	currentIndex int
	length       int
}

func isPlaindrome(bs []byte) bool {
	lastIndex := len(bs) - 1
	half := len(bs) / 2
	for i := 0; i < half; i++ {
		if bs[i] != bs[lastIndex-i] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {

	var ss string

	bs := []byte(s)
	if len(bs) > 0 {
		maxCharPos := charPos{}
		posMap := make(map[byte][]int, 26)

		for pos, v := range bs {
			if points, ok := posMap[v]; ok {

				if pos-posMap[v][0] < maxCharPos.length {
					continue
				}

				for _, point := range points {
					if pos-point < maxCharPos.length {
						continue
					}
					if isPlaindrome(bs[point : pos+1]) {
						if pos-point > maxCharPos.length {
							maxCharPos.currentIndex = pos
							maxCharPos.lastIndex = point
							maxCharPos.length = pos - point
						}
					}
				}

				posMap[v] = append(posMap[v], pos)
			} else {
				posMap[v] = []int{pos}
			}
		}
		if maxCharPos.length == 0 {
			ss = string(bs[0])
		} else {
			ss = string(bs[maxCharPos.lastIndex : maxCharPos.currentIndex+1])
		}
	}

	return ss
}

func TestIsPalindrome(t *testing.T) {
	s := []byte("adada")

	res := isPlaindrome(s)
	if !res {
		t.Error("error adjust ", s)
	}

	s = []byte("abcde")
	res = isPlaindrome(s)
	if res {
		t.Error("error adjust ", s)
	}
}

func TestLongestPalindrome(t *testing.T) {
	var s, ss string
	s = "abcdbd"
	ss = longestPalindrome(s)
	if ss != "dbd" {
		t.Error("dbd != ", ss)
	}

	s = "abcdd"
	ss = longestPalindrome(s)
	if ss != "dd" {
		t.Error("dd != ", ss)
	}

	s = "aaa"
	ss = longestPalindrome(s)
	if ss != "aaa" {
		t.Error("aa != ", ss)
	}
	s = "a"
	ss = longestPalindrome(s)
	if ss != "a" {
		t.Error("a != ", ss)
	}

	s = ""
	ss = longestPalindrome(s)
	if ss != "" {
		t.Error("\\s != ", ss)
	}

	s = "babadada"
	ss = longestPalindrome(s)

	if ss != "adada" {
		t.Error("adada != ", ss)
	}
}
