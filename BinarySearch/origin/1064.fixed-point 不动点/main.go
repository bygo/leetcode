package main

// https://leetcode-cn.com/problems/fixed-point

func fixedPoint(arr []int) int {
	lo, hi := 0, len(arr)
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if arr[mid] < mid {
			lo = mid + 1
		} else if mid <= arr[mid] {
			hi = mid
		}
	}
	if lo < len(arr) && arr[lo] == lo {
		return lo
	}
	return -1
}
