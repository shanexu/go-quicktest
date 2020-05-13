package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{1, 2, 2}))
}

func singleNumber(nums []int) int {
	x := nums[0]
	for i := 1; i < len(nums); i++ {
		x ^= nums[i]
	}
	return x
}
