package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	resMap := make(map[string]bool)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		s := make(map[int]bool)
		for j := i + 1; j < len(nums); j++ {
			x := -(nums[i] + nums[j])
			if s[x] {
				triple := []int{x, nums[i], nums[j]}
				sort.Ints(triple)
				key := tripleToKey(triple)
				if !resMap[key] {
					res = append(res, triple)
					resMap[key] = true
				}
			} else {
				s[nums[j]] = true
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
