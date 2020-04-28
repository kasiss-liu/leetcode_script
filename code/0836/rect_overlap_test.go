package main

import (
	"testing"
)

func isRectangleOverlap(rec1 []int, rec2 []int) bool {

	rec1X1 := rec1[0]
	rec1Y1 := rec1[1]
	rec1X2 := rec1[2]
	rec1Y2 := rec1[3]

	rec2X1 := rec2[0]
	rec2Y1 := rec2[1]
	rec2X2 := rec2[2]
	rec2Y2 := rec2[3]

	if rec1X1 >= rec2X2 || rec1Y1 >= rec2Y2 || rec1X2 <= rec2X1 || rec1Y2 <= rec2Y1 {
		return false
	}

	return true
}

func TestRectIsOverLap(t *testing.T) {
	var rec1, rec2 []int

	rec1 = []int{7, 8, 13, 15}
	rec2 = []int{10, 8, 12, 20}

	var exp, res bool
	exp = true
	res = isRectangleOverlap(rec1, rec2)

	if exp != res {
		t.Error("error")
	}
}
