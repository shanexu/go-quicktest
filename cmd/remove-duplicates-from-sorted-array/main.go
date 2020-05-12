package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1,1,2}))
	fmt.Println(removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4}))
}

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return n
	}
	p := nums[0]
	m := 1
	for i := 1; i < n; i++ {
		x := nums[i]
		if x != p {
			p = x
			nums[m]=p
			m++
		}
	}
	return m
}
