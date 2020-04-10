package main

import (
	"testing"
)

func reverseWords(s string) string {
	sl := len(s)

	if sl < 1 {
		return s
	}

	firstChar := false
	words := make([][]byte, 0, sl)

	word := make([]byte, 0, 5)
	for _, c := range s {
		if c == ' ' {
			if len(word) > 0 {
				words = append(words, word)
				word = make([]byte, 0, 5)
			}
			continue
		}
		if !firstChar {
			firstChar = true
		}
		word = append(word, byte(c))
	}
	if len(word) > 0 {
		words = append(words, word)
	}
	if len(words) < 1 {
		return ""
	}
	bs := make([]byte, 0, sl)
	for i := len(words) - 1; i >= 0; i-- {
		bs = append(bs, words[i]...)
		bs = append(bs, ' ')
	}
	return string(bs[:len(bs)-1])
}

func TestReversWords(t *testing.T) {
	t.Error(reverseWords("ni hao"))
}
