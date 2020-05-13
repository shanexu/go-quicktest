package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	resMap := make(map[string]bool)
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-1; i++ {
		l, r := i+1, n-1
		x := nums[i]
		for l < r {
			s := x + nums[l] + nums[r]
			if s == 0 {
				triple := sortTriple(x, nums[l], nums[r])
				key := tripleToKey(triple)
				if !resMap[key] {
					res = append(res, triple)
					resMap[key] = true
				}
				l++
				r--
			} else if s < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sortTriple(a, b, c int) []int {
	m1 := min(min(a, b), c)
	m2 := max(max(a, b), c)
	return []int{m1, a + b + c - m1 - m2, m2}
}

func tripleToKey(vs []int) string {
	return fmt.Sprintf("%d:%d:%d", vs[0], vs[1], vs[2])
}
