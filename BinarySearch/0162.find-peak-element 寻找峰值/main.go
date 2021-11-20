package main

// https://leetcode-cn.com/problems/find-peak-element

func findPeakElement(nums []int) int {
	n := len(nums)

	get := func(i int) int {
		if i == -1 || i == n {
			return -1 << 63
		}
		return nums[i]
	}

	l, r := 0, n
	for {
		mid := l + (r-l)/2
		cur := get(mid)
		if cur < get(mid+1) {
			l = mid + 1
		} else if cur < get(mid-1) {
			r = mid
		} else {
			return mid
		}
	}
}
