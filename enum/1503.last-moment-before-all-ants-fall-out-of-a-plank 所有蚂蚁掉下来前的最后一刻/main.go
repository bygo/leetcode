package main

// https://leetcode-cn.com/problems/last-moment-before-all-ants-fall-out-of-a-plank

func getLastMoment(n int, left []int, right []int) int {
	var max int
	for _, v := range left {
		if max < v {
			max = v
		}
	}

	for _, v := range right {
		if max < n-v {
			max = n - v
		}
	}

	return max
}
