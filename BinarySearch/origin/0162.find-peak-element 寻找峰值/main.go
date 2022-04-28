package main

// https://leetcode-cn.com/problems/find-peak-element

func findPeakElement(nums []int) int {
	numL := len(nums)

	find := func(i int) int {
		if i == -1 || i == numL {
			return -1 << 63
		}
		return nums[i]
	}

	lo, hi := 0, numL
	for {
		mid := int(uint(lo+hi) >> 1)
		cur := find(mid)
		if cur < find(mid+1) {
			lo = mid + 1
		} else if cur < find(mid-1) {
			hi = mid
		} else {
			return mid
		}
	}
}
