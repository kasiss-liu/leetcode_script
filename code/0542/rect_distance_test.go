package main

import "testing"

//直接在原矩阵上修改
//设起始 level = 0
//1 遍历矩阵遇到0就把4周的1设置为 level - 1
//2 level--
//3 继续第二次遍历 直到没有新level

func updateMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return nil
	}

	m, n := len(matrix), len(matrix[0])
	var helper func(i, j int)
	level := 0
	helper = func(i, j int) {
		if i < 0 || i > m-1 || j < 0 || j > n-1 {
			return
		}
		if matrix[i][j] == 1 {
			matrix[i][j] = level - 1
		}
	}
	// 遍历直到hasLv为false
	hasLv := true
	for hasLv {
		hasLv = false
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if matrix[i][j] == level {
					hasLv = true
					helper(i-1, j)
					helper(i+1, j)
					helper(i, j+1)
					helper(i, j-1)
				}
			}
		}
		level--
	}
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = -matrix[i][j]
		}
	}

	return matrix
}

func TestUpdateMatrix(t *testing.T) {
	a := [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{1, 1, 1},
	}
	b := updateMatrix(a)
	t.Error(b)
}
