package main

import "testing"

func validPalindrome(s string) bool {
	sLen := len(s)
	lastIndex := sLen - 1
	jumpNum := 1
	i := 0
	isPalindrome := true
	for i <= lastIndex-i {
		if s[i] != s[lastIndex-i] {
			if jumpNum > 0 {
				lastIndex = lastIndex - 1
				jumpNum--
				continue
			}
			isPalindrome = false
			break
		}
		i++
	}

	if isPalindrome {
		return true
	}

	isPalindrome = true
	lastIndex = sLen - 1
	j := lastIndex
	jumpNum = 1
	for j >= lastIndex-j {
		if s[j] != s[lastIndex-j] {
			if jumpNum > 0 {
				lastIndex = lastIndex + 1
				jumpNum--
				continue
			}
			isPalindrome = false
			break
		}
		j--
	}
	return isPalindrome

}

func TestValidPalindrome(t *testing.T) {
	a := "aba"

	exp := true

	ret := validPalindrome(a)

	if exp != ret {
		t.Error(exp, "not equal", ret)
	}

}
