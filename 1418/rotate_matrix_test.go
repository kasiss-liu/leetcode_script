package main

import "testing"

func rotate(matrix [][]int) [][]int {
	l := len(matrix)
	tmp := make([][]int, l)
	for n := 0; n < l; n++ {
		tmp[n] = make([]int, len(matrix[n]))
	}
	ll := l - 1
	for i := 0; i < l; i++ {
		lj := len(matrix[i])
		for j := 0; j < lj; j++ {
			tmp[j][ll-i] = matrix[i][j]
		}
	}
	copy(matrix, tmp)
	return matrix
}

func TestRotateMatrix(t *testing.T) {
	var matrix = [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	res := rotate(matrix)
	t.Error(res)
}
