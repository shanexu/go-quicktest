package main

import (
	"fmt"
	"math"
)

func main() {
	ar1 := []int64{1, 2, 3, 6}
	ar2 := []int64{4, 6, 8, 10}
	fmt.Println(getMedian(ar1, ar2))
}

func getMedian(ar1 []int64, ar2 []int64) float64 {
	n := len(ar1)
	if n <= 0 {
		return float64(-1)
	}
	if n == 1 {
		return float64(ar1[0]+ar2[0]) / 2.0
	}
	if n == 2 {
		return (math.Max(float64(ar1[0]), float64(ar2[0])) + math.Min(float64(ar1[1]), float64(ar2[1]))) / 2.0
	}
	m1 := median(ar1)
	m2 := median(ar2)

	if m1 == m2 {
		return m1
	}

	if m1 < m2 {
		if n%2 == 0 {
			return getMedian(ar1[n/2-1:], ar2[:n/2+1])
		}
		return getMedian(ar1[n/2:], ar2[:n/2+1])
	}

	if n%2 == 0 {
		return getMedian(ar1[:n/2+1], ar2[n/2-1:])
	}
	return getMedian(ar1[:n/2+1], ar2[n/2:])
}

func median(arr []int64) float64 {
	n := len(arr)
	if n%2 == 1 {
		return float64(arr[n/2])
	}
	return float64((arr[n/2-1] + arr[n/2])) / 2.0
}
