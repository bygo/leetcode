package main

import "sort"

// https://leetcode-cn.com/problems/find-the-distance-value-between-two-arrays

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2)
	var res int
	var binarySearch = func(target int) int {
		l, r := 0, len(arr2)
		for l < r {
			mid := l + (r-l)/2
			cur := abs(target - arr2[mid])
			if cur <= d {
				return 0
			}

			if arr2[mid] <= target {
				l = mid + 1
			} else {
				r = mid
			}
		}
		return 1
	}
	for i := range arr1 {
		res += binarySearch(arr1[i])
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
