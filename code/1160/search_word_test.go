package main

import (
	"testing"
)

func countCharacters(words []string, chars string) int {
	bchars := []byte(chars)
	bmap := make([]int, 256)
	for _, char := range bchars {
		i := int(char)
		bmap[i]++
	}
	wordLen := 0
	for _, word := range words {
		bword := []byte(word)
		if len(bword) > len(chars) {
			continue
		}

		cpBmap := make([]int, 256)
		copy(cpBmap, bmap)
		needSum := true
		for _, b := range bword {
			i := int(b)
			cpBmap[i]--
			if cpBmap[i] < 0 {
				needSum = false
				break
			}

		}

		if needSum {
			wordLen += len(bword)
		}
	}

	return wordLen
}

func TestCountCharacters(t *testing.T) {
	words := []string{"cat", "bt", "hat", "tree"}
	chars := "atach"

	exp := 6
	res := countCharacters(words, chars)
	if exp != res {
		t.Error(exp, " not equal ", res)
	}
}
