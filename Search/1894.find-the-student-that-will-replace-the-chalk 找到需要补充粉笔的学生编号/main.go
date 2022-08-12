package main

// https://leetcode.cn/problems/find-the-student-that-will-replace-the-chalk

func chalkReplacer(chalk []int, k int) int {
	r := len(chalk)
	for i := 1; i < r; i++ {
		chalk[i] += chalk[i-1]
	}

	k %= chalk[r-1]
	k += 1
	var l = 0
	var mid int
	for l < r {
		mid = l + (r-l)/2
		if chalk[mid] < k {
			l = mid + 1
		} else if k <= chalk[mid] {
			r = mid - 1
		}
	}
	return l
}
