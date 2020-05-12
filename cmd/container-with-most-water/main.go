package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	m := 0
	for i := 1; i < len(height); i++ {
		hi := height[i]
		for j := 0; j < i; j++ {
			s := min(height[j], hi) * (i - j)
			if s > m {
				m = s
			}
		}
	}
	return m
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
