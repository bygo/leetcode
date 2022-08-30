package main

import "sort"

// https://leetcode.cn/problems/find-the-distance-value-between-two-arrays

// general
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2)
	aL := len(arr2)
	var search = func(target int) int {
		lo, hi := 0, aL
		for lo < hi {
			mid := int(uint(lo+hi) >> 1)
			cur := abs(target - arr2[mid])
			if cur <= d { // TODO
				return 0
			}
			if arr2[mid] <= target {
				lo = mid + 1
			} else if target < arr2[mid] {
				hi = mid
			}
		}
		return 1
	}

	var cnt int
	for i := range arr1 {
		cnt += search(arr1[i])
	}
	return cnt
}

// offset
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2)
	aL := len(arr2)
	var search = func(target int) int {
		lo, hi := -1, aL-1
		for lo < hi {
			mid := int(uint(lo+hi+1) >> 1)
			cur := abs(target - arr2[mid])
			if cur <= d {
				return 0
			}

			if arr2[mid] <= target {
				lo = mid
			} else if target < arr2[mid] {
				hi = mid - 1
			}
		}
		return 1
	}
	var cnt int
	for i := range arr1 {
		cnt += search(arr1[i])
	}
	return cnt
}

// equal
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2)
	aL := len(arr2)
	var search = func(target int) int {
		lo, hi := 0, aL-1
		for lo <= hi {
			mid := int(uint(lo+hi) >> 1)
			cur := abs(target - arr2[mid])
			if cur <= d {
				return 0
			}

			if arr2[mid] <= target {
				lo = mid + 1
			} else if target <= arr2[mid] {
				hi = mid - 1
			}
		}
		return 1
	}
	var cnt int
	for i := range arr1 {
		cnt += search(arr1[i])
	}
	return cnt
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
