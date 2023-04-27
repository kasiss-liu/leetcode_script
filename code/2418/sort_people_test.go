package main

import (
	"testing"
)

func SortPeople(names []string, heights []int) []string {

	var qs func(int, int)
	qs = func(i, j int) {
		if i >= j {
			return
		}
		v := heights[i]
		l, h := i, j
		for l != h {
			for heights[h] <= v && l < h {
				h--
			}
			for heights[l] >= v && l < h {
				l++
			}
			if l < h {
				heights[l], heights[h] = heights[h], heights[l]
				names[l], names[h] = names[h], names[l]
			}
		}
		heights[l], heights[i] = heights[i], heights[l]
		names[l], names[i] = names[i], names[l]
		qs(i, l-1)
		qs(l+1, j)
	}
	qs(0, len(heights)-1)

	return names
}

func TestSortPeople(t *testing.T) {

	a := []int{4, 3, 1, 7, 2, 9}
	b := []string{"a1", "a2", "a3", "a4", "a5", "a6"}
	t.Log(SortPeople(b, a))

}
