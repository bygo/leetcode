package main

// https://leetcode.cn/problems/find-peak-element

// general
func findPeakElement(nums []int) int {
	numsL := len(nums)
	find := func(i int) int {
		if i == -1 || i == numsL {
			return -1 << 63
		}
		return nums[i]
	}

	lo, hi := 0, numsL
	for {
		mid := int(uint(lo+hi) >> 1)
		cur := find(mid)
		if cur < find(mid+1) {
			lo = mid + 1
		} else if cur < find(mid-1) {
			hi = mid - 1
		} else {
			return mid
		}
	}
}

// offset
func findPeakElement(nums []int) int {
	numsL := len(nums)
	find := func(i int) int {
		if i == -1 || i == numsL {
			return -1 << 63
		}
		return nums[i]
	}

	lo, hi := -1, numsL-1
	for {
		mid := int(uint(lo+hi+1) >> 1)
		cur := find(mid)
		if cur < find(mid+1) {
			lo = mid + 1
		} else if cur < find(mid-1) {
			hi = mid - 1
		} else {
			return mid
		}
	}
}

// equal
func findPeakElement(nums []int) int {
	numsL := len(nums)
	find := func(i int) int {
		if i == -1 || i == numsL {
			return -1 << 63
		}
		return nums[i]
	}

	lo, hi := 0, numsL-1
	for {
		mid := int(uint(lo+hi) >> 1)
		cur := find(mid)
		if cur < find(mid+1) {
			lo = mid + 1
		} else if cur < find(mid-1) {
			hi = mid - 1
		} else {
			return mid
		}
	}
}
