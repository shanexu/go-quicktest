package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMediaSortedArrays([]int{1, 2, 3}, []int{4, 5, 6, 7, 8}))
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func mo2(a, b int) float64 {
	return float64(a+b) / 2.0
}

func mo3(a, b, c int) int {
	return a + b + c - max(a, max(b, c)) - min(a, min(b, c))
}

func mo4(a, b, c, d int) float64 {
	max := max(a, max(b, max(c, d)))
	min := min(a, min(b, min(c, d)))
	return float64(a+b+c+d-max-min) / 2.0
}

func medianSingle(arr []int) float64 {
	n := len(arr)
	if n == 0 {
		return -1.0
	}
	if n%2 == 0 {
		return float64(arr[n/2]+arr[n/2-1]) / 2.0
	}
	return float64(arr[n/2])
}

func findMedianUtil(a []int, b []int) float64 {
	n := len(a)
	m := len(b)
	if n == 0 {
		return medianSingle(b)
	}
	if n == 1 {
		if m == 1 {
			return mo2(a[0], b[0])
		}
		if m&1 > 0 {
			return mo2(b[m/2], mo3(a[0], b[m/2-1], b[m/2+1]))
		}
		return float64(mo3(b[m/2], b[m/2-1], a[0]))
	} else if n == 2 {
		if m == 2 {
			return mo4(a[0], a[1], b[0], b[1])
		}
		if m&1 > 0 {
			return float64(mo3(b[m/2], max(a[0], b[m/2-1]), min(a[1], b[m/2+1])))
		}
		return mo4(b[m/2], b[m/2-1], max(a[0], b[m/2-2]), min(a[1], b[m/2+1]))
	}

	idxA := (n - 1) / 2
	idxB := (m - 1) / 2

	if a[idxA] <= b[idxB] {
		return findMedianUtil(a[idxA:], b[:m-idxA])
	}

	return findMedianUtil(a[:n-idxA], b[idxA:])
}

func findMediaSortedArrays(nums1, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)
	if n > m {
		return findMedianUtil(nums2, nums1)
	}
	return findMedianUtil(nums1, nums2)
}
