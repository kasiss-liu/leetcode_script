package main

import "testing"

func isValid(s string) bool {
	bs := []byte(s)
	stack := make([]byte, 0, len(bs)/2+1)
	for _, b := range bs {
		//压栈 ( [ {
		if b == '(' || b == '[' || b == '{' {
			stack = append(stack, b)
		}
		//出栈 ) ] }
		if b == ')' || b == ']' || b == '}' {
			if len(stack) > 0 {
				c := stack[len(stack)-1]
				if checkEqual(b, c) {
					stack = stack[:len(stack)-1]
					continue
				}
			}
			return false
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}

func checkEqual(b, c byte) (res bool) {
	res = true
	switch b {
	case ')':
		if c != '(' {
			res = false
		}
	case ']':
		if c != '[' {
			res = false
		}
	case '}':
		if c != '{' {
			res = false
		}
	default:
		res = false
	}
	return
}

func TestParenthesisValid(t *testing.T) {
	var s string
	var exp, res bool

	s = "()"
	exp = true
	res = isValid(s)
	if exp != res {
		t.Error(exp, "not equal", res)
	}

	s = "()[]{}"
	exp = true
	res = isValid(s)
	if exp != res {
		t.Error(exp, "not equal", res)
	}

	s = "([)]"
	exp = false
	res = isValid(s)
	if exp != res {
		t.Error(exp, "not equal", res)
	}
	s = ""
	exp = true
	res = isValid(s)
	if exp != res {
		t.Error(exp, "not equal", res)
	}

	s = "{[]}"
	exp = true
	res = isValid(s)
	if exp != res {
		t.Error(exp, "not equal", res)
	}

	s = "}"
	exp = false
	res = isValid(s)
	if exp != res {
		t.Error(exp, "not equal", res)
	}
}
