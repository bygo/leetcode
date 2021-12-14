package main

// https://leetcode-cn.com/problems/teemo-attacking

func findPoisonedDuration(timeSeries []int, duration int) int {
	var cnt int
	l := len(timeSeries)
	top := l - 1
	cnt = l * duration // 最多中毒时间
	for i := top; 0 < i; i-- {
		cur := timeSeries[i] - timeSeries[i-1] // 差值
		if cur < duration {
			cnt -= duration
			cnt += cur
		}
	}
	return cnt
}
