package main

import "testing"

func longestPalindrome(s string) int {
	charMap := [123]int{}
	evenLength := 0
	oddLength := 0
	for _, b := range s {
		i := int(b)
		charMap[i]++
		if charMap[i]%2 == 0 {
			evenLength += 2
		}
	}

	for _, chr := range charMap {
		if chr > 0 && chr%2 != 0 {
			oddLength = 1
			break
		}
	}
	return oddLength + evenLength
}

func TestLongestPalindrome(t *testing.T) {
	s := "abccccdd"

	var exp, res int
	exp = 7
	res = longestPalindrome(s)
	if exp != res {
		t.Error("expected ", exp, "got ", res)
	}
}
