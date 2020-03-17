package main

import "testing"

func countCharacters(words []string, chars string) int {
	bchars := []byte(chars)
	bmap := make(map[byte]int, len(bchars))
	for _, char := range bchars {
		if _, ok := bmap[char]; ok {
			bmap[char]++
		} else {
			bmap[char] = 1
		}
	}
	wordLen := 0
	for _, word := range words {
		bword := []byte(word)
		if len(bword) > len(chars) {
			continue
		}
		wordMap := make(map[byte]int, len(bword))
		needSum := true
		for _, b := range bword {
			if _, ok := bmap[b]; !ok {
				needSum = false
				break
			} else {
				if _, ok := wordMap[b]; ok {
					wordMap[b]++
				} else {
					wordMap[b] = 1
				}
				if wordMap[b] > bmap[b] {
					needSum = false
					break
				}
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
