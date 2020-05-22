package main

import "testing"

func isPalindrome(s string) bool {

	sLen := len(s)
	i := 0
	j := sLen - 1

	for i <= j {
		if !isLetter(byte(s[i])) {
			i++
			continue
		}
		if !isLetter(byte(s[j])) {
			j--
			continue
		}
		if toLower(byte(s[i])) != toLower(byte(s[j])) {
			return false
		}
		i++
		j--
	}
	return true
}
func isLetter(b byte) bool {
	if '0' <= b && b <= '9' {
		return true
	}
	if 'a' <= b && b <= 'z' {
		return true
	}
	if 'A' <= b && b <= 'Z' {
		return true
	}
	return false
}

func toLower(b byte) byte {
	if '0' <= b && b <= '9' {
		return b
	}
	if 'a' <= b && b <= 'z' {
		return b
	}
	return b + ('a' - 'A')
}

func TestIsPalindrome(t *testing.T) {
	a := "abccba"
	exp := true

	ret := isPalindrome(a)
	if exp != ret {
		t.Error(exp, "not equal", ret)
	}
}
