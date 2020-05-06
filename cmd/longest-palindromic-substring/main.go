package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("aacdefcaa"))
}

func longestPalindrome(s string) string {
	l := len(s)
	if l == 0 {
		return ""
	}
	maxLength := 1
	var (
		start int
		low   int
		high  int
	)
	for i := 1; i < l; i++ {
		low = i - 1
		high = i
		for low >= 0 && high < l && s[low] == s[high] {
			if high-low+1 > maxLength {
				start = low
				maxLength = high - low + 1
			}
			low--
			high++
		}

		low = i - 1
		high = i + 1
		for low >= 0 && high < l && s[low] == s[high] {
			if high-low+1 > maxLength {
				start = low
				maxLength = high - low + 1
			}
			low--
			high++
		}
	}
	return s[start : start+maxLength]
}
