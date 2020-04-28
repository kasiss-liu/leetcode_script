package main

import (
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	l := 0
	bs := []byte(s)
	strlen := len(bs)
	chars := make([]byte, 0, 10)

	for i := 0; i < strlen; i++ {
		b := bs[i]

		for key, char := range chars {
			if char == b {
				chars = chars[key+1:]
				break
			}
		}
		chars = append(chars, b)
		if l < len(chars) {
			l = len(chars)
		}

	}
	return l
}

func TestLongestSubstring(t *testing.T) {
	var len int
	var s string
	s = "pwwkew"
	len = lengthOfLongestSubstring(s)
	if len != 3 {
		t.Error("test longest substring failed expected ", 3, " get ", len)
	}

	s = "abcabcbb"
	len = lengthOfLongestSubstring(s)
	if len != 3 {
		t.Error("test longest substring failed expected ", 3, " get ", len)
	}

	s = "abcabcbddsgadsfagasdgasdfb"
	len = lengthOfLongestSubstring(s)
	if len != 6 {
		t.Error("test longest substring failed expected ", 6, " get ", len)
	}
}
