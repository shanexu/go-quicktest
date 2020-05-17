package main

import "fmt"

func main() {
	// TODO
	fmt.Println(minSwapsCouples([]int{28,4,37,54,35,41,43,42,45,38,19,51,49,17,47,25,12,53,57,20,2,1,9,27,31,55,32,48,59,15,14,8,3,7,58,23,10,52,22,30,6,21,24,16,46,5,33,56,18,50,39,34,29,36,26,40,44,0,11,13}))
}

func minSwapsCouples(row []int) int {
	index := make([]int, len(row), len(row))
	for i := range row {
		index[row[i]] = i
	}
	return minSwapsCouplesUtil(row, index, 0, len(row))
}

func minSwapsCouplesUtil(arr []int, index []int, i, n int) int {
	if i > n-1 {
		return 0
	}
	if pair(arr[i]) == arr[i+1] {
		return minSwapsCouplesUtil(arr, index, i+2, n)
	}
	one := arr[i+1]
	indexTwo := i + 1
	indexOne := index[pair(arr[i])]
	two := arr[index[pair(arr[i])]]
	arr[i+1], arr[indexOne] = arr[indexOne], arr[i+1]
	updateIndex(index, one, indexOne, two, indexTwo)
	a := minSwapsCouplesUtil(arr, index, i+2, n)

	arr[i+1], arr[indexOne] = arr[indexOne], arr[i+1]
	updateIndex(index, one, indexTwo, two, indexOne)
	one = arr[i]
	indexOne = index[pair(arr[i+1])]
	two = arr[index[pair(arr[i+1])]]
	indexTwo = i
	arr[i], arr[indexOne] = arr[indexOne], arr[i]
	updateIndex(index, one, indexOne, two, indexTwo)
	b := minSwapsCouplesUtil(arr, index, i+2, n)

	arr[i], arr[indexOne] = arr[indexOne], arr[i]
	updateIndex(index, one, indexTwo, two, indexOne)
	if a < b {
		return 1 + a
	}
	return 1 + b
}

func pair(p int) int {
	if p/2*2 == p {
		return p/2*2 + 1
	}
	return p / 2 * 2
}

func updateIndex(index []int, a, ai, b, bi int) {
	index[a] = ai
	index[b] = bi
}
