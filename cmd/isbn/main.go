package main

import (
	"fmt"
)

func main() {
	fmt.Println(ValidISBN10("1112223339"))
}

func ValidISBN10(isbn string) bool {
	if len(isbn) != 10 {
		return false
	}
	sum := 0
	for i, c := range isbn {
		if (c >= '0' && c <= '9') || (c == 'X' && i == 9) {
			v := int(c - '0')
			if c == 'X' {
				v = 10
			}
			sum += v * (i + 1)
		} else {
			return false
		}
	}
	return sum%11 == 0
}
