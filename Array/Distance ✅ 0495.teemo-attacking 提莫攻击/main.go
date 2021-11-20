package main

// https://leetcode-cn.com/problems/teemo-attacking

func findPoisonedDuration(timeSeries []int, duration int) int {
	var res int
	l := len(timeSeries)
	top := l - 1
	res = l * duration
	for i := top; 0 < i; i-- {
		cur := timeSeries[i] - timeSeries[i-1]
		if cur < duration {
			res -= duration - cur
		}
	}
	return res
}
