package main

// https://leetcode-cn.com/problems/first-bad-version

func firstBadVersion(r int) int {
	var l = 1
	var mid int
	for l <= r {
		mid = l + (r-l)/2
		if isBadVersion(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func isBadVersion(n int) bool {
	return n == 1
}
