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
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}

	t.Error(merge(a))
}

func TestSortArr(t *testing.T) {
	a := [][]int{
		{4, 5},
		{1, 4},
		{0, 1},
	}
	t.Log(sortArr(a))
}

func sortArr(arr [][]int) [][]int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i][0] > arr[j][0] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func TestMergeArr(t *testing.T) {
	a := [][]int{
		{4, 5},
		{1, 4},
		{0, 1},
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})
	t.Log(a)
}

func mergeArr(intervals [][]int) [][]int {
	res := make([][]int, 0, 10)
	if len(intervals) < 2 {
		return intervals
	}
	intervals = sortArr(intervals)
	tmp := intervals[0]
	for i := 1; i < len(intervals); i++ {
		arr := intervals[i]
		if arr[0] > tmp[1] {
			res = append(res, tmp)
			tmp = arr
			if i == len(intervals)-1 {
				res = append(res, tmp)
			}
			continue
		}
		if arr[1] > tmp[1] {
			tmp[1] = arr[1]
		}
		if i == len(intervals)-1 {
			res = append(res, tmp)
		}

	}
	return res
}

func TestSortArrs(t *testing.T) {
	a := []int{4, 1, 0, 9, 5, 8, 7, 6}
	t.Log(sortArrs(a))
}

func sortArrs(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
