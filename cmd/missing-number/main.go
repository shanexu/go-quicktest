package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{0, 1, 2}))
}

func missingNumber(nums []int) int {
	sum := 0
	for i := 1; i <= len(nums); i++ {
		sum += i
	}
	for _, t := range nums {
		sum -= t
	}
	return sum
}
