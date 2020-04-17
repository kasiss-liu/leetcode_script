package main

import (
	"sort"
	"testing"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := [][]int{}
	result = append(result, intervals[0])

	for i := 1; i < len(intervals); i++ {
		m := result[len(result)-1]
		c := intervals[i]

		if c[0] > m[1] {
			result = append(result, c)
			continue
		}

		if c[1] > m[1] {
			m[1] = c[1]
		}
	}

	return result

}

func TestMerge(t *testing.T) {
	a := [][]int{
		[]int{1, 3},
		[]int{2, 6},
		[]int{8, 10},
		[]int{15, 18},
	}

	t.Error(merge(a))
}
