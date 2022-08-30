package main

import "sort"

// https://leetcode.cn/problems/first-bad-version

func firstBadVersion(hi int) int {
	return sort.Search(hi+1, isBadVersion) // TODO
}

func isBadVersion(n int) bool {
	return n == 1
}
