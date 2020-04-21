package main

import "testing"

func strStr(haystack string, needle string) int {
	lh, ln := len(haystack), len(needle)

	if ln == 0 {
		return 0
	}

	if lh < ln {
		return -1
	}

	for i := 0; i <= lh-ln; i++ {
		tmp := haystack[i : i+ln]
		if string(tmp) == needle {
			return i
		}
	}
	return -1
}

func TestStrstr(t *testing.T) {
	haystack := "aaaaa"
	needle := "bba"

	exp := -1
	res := strStr(haystack, needle)
	if exp != res {
		t.Error(exp, "not equal", res)
	}
}
