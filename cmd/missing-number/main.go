package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{0,1,2}))
}

func missingNumber(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		nums[i] += 1
	}
	for i := 0; i < n; i++ {
		if abs(nums[i])-1 < n {
			nums[abs(nums[i])-1] = -nums[abs(nums[i])-1]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i
		}
	}
	return n
}

func abs(v int) int {
	switch {
	case v < 0:
		return -v
	default:
		return v
	}
}
