package main

// https://leetcode-cn.com/problems/peak-index-in-a-mountain-array

func peakIndexInMountainArray(arr []int) int {
	lo, hi := 0, len(arr)
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if arr[mid] < arr[mid+1] {
			lo = mid + 1
		} else if arr[mid] < arr[mid-1] {
			hi = mid
		} else {
			return mid
		}
	}
	return -1
}
