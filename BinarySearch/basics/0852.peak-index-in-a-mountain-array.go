package main

// https://leetcode.cn/problems/peak-index-in-a-mountain-array

// general
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

// offset
func peakIndexInMountainArray(arr []int) int {
	lo, hi := -1, len(arr)-1
	for lo < hi {
		mid := int(uint(lo+hi+1) >> 1)
		if arr[mid] < arr[mid+1] {
			lo = mid
		} else if arr[mid] < arr[mid-1] {
			hi = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

// equal
func peakIndexInMountainArray(arr []int) int {
	lo, hi := 1, len(arr)-2
	for lo <= hi {
		mid := int(uint(lo+hi) >> 1)
		if arr[mid] < arr[mid+1] {
			lo = mid + 1
		} else if arr[mid] < arr[mid-1] {
			hi = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
