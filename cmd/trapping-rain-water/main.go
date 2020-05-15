package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	n := len(height)
	left := make([]int, n, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			left[i] = height[i]
			continue
		}
		left[i] = max(left[i-1], height[i])
	}
	right := make([]int, n, n)
	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			right[i] = height[i]
			continue
		}
		right[i] = max(right[i+1], height[i])
	}
	s := 0
	for i := 0; i < n; i++ {
		h := height[i]
		l := left[i]
		r := right[i]
		if h < l && h < r {
			s += min(l, r) - h
		}
	}
	return s
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
