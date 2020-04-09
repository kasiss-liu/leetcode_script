package main

import (
	"reflect"
	"testing"
)

func maxDepthAfterSplit(seq string) []int {
	stack := make([]int, 0, len(seq))
	var deep = -1
	for _, v := range seq {
		if string(v) == "(" {
			deep++
			stack = append(stack, deep%2)
		}
		if string(v) == ")" {
			stack = append(stack, deep%2)
			deep--
		}
	}
	return stack
}

func TestDepth(t *testing.T) {
	seq := "(()())"
	exp := []int{0, 1, 1, 1, 1, 0}

	res := maxDepthAfterSplit(seq)
	if !reflect.DeepEqual(exp, res) {
		t.Error(exp, "not equal", res)
	}
}
