package main

// https://leetcode-cn.com/problems/peak-index-in-a-mountain-array

func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)
	for l < r {
		mid := l + (r-l)/2
		if arr[mid] < arr[mid+1] {
			l = mid + 1
		} else if arr[mid] < arr[mid-1] {
			r = mid
		} else {
			return mid
		}
	}
	return -1
}
