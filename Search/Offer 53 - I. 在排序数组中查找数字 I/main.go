package main

// https://leetcode.cn/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
func search(nums []int, target int) int {
	return s(nums, target+1) - s(nums, target)
}

func s(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)/2
		cur := nums[mid]
		if target <= cur {
			r = mid
		} else if cur < target {
			l = mid + 1
		}
	}
	return r
}
