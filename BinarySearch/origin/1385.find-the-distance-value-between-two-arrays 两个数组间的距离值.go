package main

import "sort"

// https://leetcode-cn.com/problems/find-the-distance-value-between-two-arrays

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2)
	var binarySearch = func(target int) int {
		lo, hi := 0, len(arr2)
		for lo < hi {
			mid := int(uint(lo+hi) >> 1)
			cur := abs(target - arr2[mid])
			if cur <= d {
				return 0
			}

			if arr2[mid] <= target {
				// 比target 小，接近 target
				lo = mid + 1
			} else if target < arr2[mid] {
				// 比target 大，接近 target
				hi = mid
			}
		}
		return 1
	}

	var cnt int
	for i := range arr1 {
		cnt += binarySearch(arr1[i])
	}
	return cnt
}

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2)
	aL := len(arr2)
	var search = func(target int) int {
		lo, hi := 0, aL
		for lo < hi {
			mid := int(uint(lo+hi) >> 1)
			cur := abs(target - arr2[mid])
			if cur <= d {
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

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	aL := len(arr2)
	sort.Ints(arr2)
	var search = func(target int) int {
		lo, hi := 0, aL
		for lo < hi {
			mid := int(uint(lo+hi) >> 1)
			cur := abs(target - arr2[mid])
			if cur <= d {
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
	var cnt = 0
	for i := range arr1 {
		cnt += search(arr1[i])
	}
	return cnt
}

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	aL := len(arr2)
	sort.Ints(arr2)
	var search = func(target int) int {
		lo, hi := 0, aL
		for lo < hi {
			mid := int(uint(lo+hi) >> 1)
			cur := abs(target - arr2[mid])
			if cur <= d {
				return 0
			}
			if arr2[mid] == target {
				return 1
			} else if arr2[mid] < target {
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

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	aL := len(arr2)
	sort.Ints(arr2)
	var search = func(target int) int {
		lo, hi := 0, aL
		for lo < hi {
			mid := int(uint(lo+hi) >> 1)
			cur := abs(target - arr2[mid])
			if cur <= d {
				return 0
			}
			if arr2[mid] == target {
				return target
			} else if arr2[mid] < target {
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
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
