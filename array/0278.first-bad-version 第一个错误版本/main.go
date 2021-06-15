package main

//Link: https://leetcode-cn.com/problems/first-bad-version

func firstBadVersion(r int) int {
	var l = 1
	var mid int
	for l < r {
		mid = l + (r-l)/2
		if isBadVersion(mid) {
			r = mid // 也可能是 答案 不需要-1
		} else {
			l = mid + 1
		}
	}
	return l
}

func isBadVersion(n int) bool {
	return true
}
