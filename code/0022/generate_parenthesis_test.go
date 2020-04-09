//深度优先搜索算法
package main

import (
	"reflect"
	"testing"
)

func generateParenthesis(n int) []string {
	s := make([]string, 0, n)
	gen("", 0, 0, n, &s)
	return s
}

func gen(cur string, left, right, n int, s *[]string) {
	//右括号不能比左括号先出现
	if left < right {
		return
	}
	//如果括号数量足够 则停止
	if left == n && right == n {
		*s = append(*s, cur)
		return
	}
	//先增加左括号
	if left < n {
		gen(cur+"(", left+1, right, n, s)
	}
	//再增加右括号
	if right < n {
		gen(cur+")", left, right+1, n, s)
	}
}

func TestGenerateParenthesis(t *testing.T) {
	var n int
	var exp, res []string
	n = 3

	exp = []string{"((()))", "(()())", "(())()", "()(())", "()()()"}

	res = generateParenthesis(n)
	if !reflect.DeepEqual(exp, res) {
		t.Error(exp, "not equal", res)
	}

}
