package main

import "fmt"

func main() {
	fmt.Println(minSwapsCouples([]int{0, 2, 1, 3}))
	fmt.Println(minSwapsCouples([]int{3, 2, 0, 1}))
}

func minSwapsCouples(row []int) int {
	// toMap
	index := make([]int, len(row), len(row))
	for i := range row {
		index[row[i]] = i
	}
	fmt.Println(row)
	fmt.Println(index)
	return 0
}

func pair(p int) int {
	if p/2*2 == p {
		return p/2*2 + 1
	}
	return p / 2 * 2
}
