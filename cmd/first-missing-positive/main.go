package main

import "fmt"

func main() {
	fmt.Println(firstMissingPositive([]int{1,2,3,100}))
}

func firstMissingPositive(nums []int) int {
	arr := segregate(nums)
	return findMissingPositive(arr)
}

func segregate(arr []int) []int {
	var i, j int
	for i = 0; i < len(arr); i++ {
		if arr[i] <= 0 {
			arr[i], arr[j] = arr[j], arr[i]
			j++
		}
	}
	return arr[j:]
}

func abs(x int) int {
	switch {
	case x < 0:
		return -x
	default:
		return x
	}
}

// 负数用于标识，小于size的整数没有缺失，所以第一个缺失的整数就是第一个为正数的位置
func findMissingPositive(arr []int) int {
	size := len(arr)
	for i := 0; i < size; i++ {
		if abs(arr[i])-1 < size && arr[abs(arr[i])-1] > 0 {
			arr[abs(arr[i])-1] = -arr[abs(arr[i])-1]
		}
	}
	for i := 0; i < size; i++ {
		if arr[i] > 0 {
			return i + 1
		}
	}
	return size + 1
}
