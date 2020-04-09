package main

import "testing"

func isValid(s string) bool {
	bs := []byte(s)
	stack := make([]byte, 0, len(bs)/2+1)
	for _, b := range bs {
		//压栈 ( [ {
		if b == 40 || b == 91 || b == 123 {
			stack = append(stack, b)
		}
		//出栈 ) ] }
		if b == 41 || b == 93 || b == 125 {
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
	//)
	case 41:
		if c != 40 {
			res = false
		}
	//]
	case 93:
		if c != 91 {
			res = false
		}
	//}
	case 125:
		if c != 123 {
			res = false
		}
	//unexpected char
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
