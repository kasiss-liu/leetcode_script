package main

import (
	"reflect"
	"testing"
)

func getLeastNumbers(arr []int, k int) []int {
	res := quickSort(arr)
	return res[:k]
}

func bubbleSort(arr []int) []int {
	l := len(arr)
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func quickSort(arr []int) []int {
	l := len(arr)
	if l < 2 {
		return arr
	}

	c := arr[0]
	left, right := make([]int, 0, l), make([]int, 0, l)

	for i := 1; i < l; i++ {
		if arr[i] < c {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	left = quickSort(left)
	right = quickSort(right)

	left = append(left, c)
	return append(left, right...)
}

func TestSort(t *testing.T) {
	var arr, exp, res []int
	var k int

	arr = []int{0, 0, 1, 2, 4, 2, 2, 3, 1, 4}
	k = 8
	exp = []int{0, 0, 1, 1, 2, 2, 2, 3}
	res = getLeastNumbers(arr, k)
	if !reflect.DeepEqual(exp, res) {
		t.Error(exp, "not equal", res)
	}
}
