package main

import "fmt"

func main() {
	fmt.Println(convert("LEETCODEISHIRING", 4))
}

func convert(str string, row int) string {
	if row <= 1 {
		return str
	}
	s := []rune(str)
	r := make([]rune, len(s))
	t := 2*row - 2
	n := len(s)
	j := 0
	for i := 0; i < n; i += t {
		r[j] = s[i]
		j++
	}

	for k := 1; k < row-1; k++ {
		for i := k; i < n; i += t {
			r[j] = s[i]
			j++
			ii := i - k + (t - k)
			if ii < n {
				r[j] = s[ii]
				j++
			}
		}
	}

	for i := row - 1; i < n; i += t {
		r[j] = s[i]
		j++
	}

	return string(r)
}
