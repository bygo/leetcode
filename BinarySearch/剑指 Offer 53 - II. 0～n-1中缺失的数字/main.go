package main

func missingNumber(nums []int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] <= mid {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
