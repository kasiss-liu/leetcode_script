package main

import "testing"

func game(guess []int, answer []int) int {
	n := 0
	if len(guess) == len(answer) {
		for i := 0; i < len(guess); i++ {
			if guess[i] == answer[i] {
				n++
			}
		}
	}
	return n
}

func TestGuess(t *testing.T) {
	guess := []int{1, 2, 3}
	answer := []int{1, 2, 3}
	n := game(guess, answer)

	if n != 3 {
		t.Error("wrong guess result")
	}
}
