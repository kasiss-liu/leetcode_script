package main

import (
	"math"
	"testing"
)

func getMatrixposition(i, numRows, capNum int) (pos [2]int) {

	if numRows <= 1 {
		pos[0] = i
		pos[1] = 0
		return
	}

	var v, h int
	//计算水平偏移坐标
	//n 满单位矩阵个数
	n := int(math.Floor(float64(i) / float64(capNum)))
	//水平偏移起始位置
	h = n * (numRows - 1)
	//有多余的元素 单独计算元素位置
	m := i - n*capNum + 1

	if m > numRows {
		h += (m - numRows)
	}
	//计算垂直坐标
	if m <= numRows {
		v = m - 1
	} else {
		v = 2*numRows - m - 1
	}

	pos[0] = h
	pos[1] = v

	return
}

func convertByMatrix(s string, numRows int) string {

	bs := []byte(s)
	l := len(bs)

	var matrix [][]byte
	matrix = make([][]byte, numRows, numRows)

	capNum := 2 * (numRows - 1)

	width := l
	if numRows > 1 {
		width = int(math.Ceil(float64(l)/float64(capNum))) * (numRows - 1)
	}

	for k := 0; k < numRows; k++ {
		matrix[k] = make([]byte, width, width)
	}

	for i := 0; i < l; i++ {
		pos := getMatrixposition(i, numRows, capNum)
		matrix[pos[1]][pos[0]] = bs[i]
	}
	buf := make([]byte, 0, l)
	for _, row := range matrix {
		for _, b := range row {
			if b > 0 {
				buf = append(buf, b)
			}
		}
	}
	return string(buf)
}

func TestConvertByMatrix(t *testing.T) {

	exp := "acebdf"
	s := "abcdef"
	ss := convertByMatrix(s, 1)
	t.Log(ss)
	if exp != ss {
		t.Error(exp, " not Equal ", ss)
	}
}
