package main

import "fmt"

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("{[]}"))
	fmt.Println(isValid("]"))
}

func isValid(s string) bool {
	rs := []rune(s)
	st := make([]rune, 0, 16)

	push := func(r rune) {
		st = append(st, r)
	}
	pop := func() rune {
		if len(st) == 0 {
			return 0
		}
		n := len(st)
		r := st[n-1]
		st = st[0 : n-1]
		return r
	}

	for _, r := range rs {
		switch r {
		case ')':
			if pop() != '(' {
				return false
			}
		case '}':
			if pop() != '{' {
				return false
			}
		case ']':
			if pop() != '[' {
				return false
			}
		default:
			push(r)
		}
	}
	return len(st) == 0
}
