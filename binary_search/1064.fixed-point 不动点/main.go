package main

// https://leetcode-cn.com/problems/fixed-point

func fixedPoint(arr []int) int {
	l, r := 0, len(arr)
	for l < r {
		mid := l + (r-l)/2
		if mid <= arr[mid] {
			r = mid
		} else if arr[mid] < mid {
			l = mid + 1
		}
	}
	l--
	if 0 <= l && arr[l] == l {
		return l
	}
	return -1
}
