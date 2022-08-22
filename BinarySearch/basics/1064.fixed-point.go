package main

// https://leetcode.cn/problems/fixed-point

// general
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

// offset
func fixedPoint(arr []int) int {
	lo, hi := -1, len(arr)-1
	for lo < hi {
		mid := int(uint(lo+hi+1) >> 1)
		if arr[mid] < mid {
			lo = mid
		} else if mid <= arr[mid] {
			hi = mid - 1
		}
	}
	lo += 1
	if lo < len(arr) && arr[lo] == lo {
		return lo
	}
	return -1
}

// equal
func fixedPoint(arr []int) int {
	lo, hi := 0, len(arr)-1
	for lo <= hi {
		mid := int(uint(lo+hi) >> 1)
		if arr[mid] < mid {
			lo = mid + 1
		} else if mid <= arr[mid] {
			hi = mid - 1
		}
	}
	if lo < len(arr) && arr[lo] == lo {
		return lo
	}
	return -1
}
