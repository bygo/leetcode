package main

// https://leetcode.cn/problems/teemo-attacking

func findPoisonedDuration(timeSeries []int, duration int) int {
	var cntMax int
	timeL := len(timeSeries)
	top := timeL - 1
	cntMax = timeL * duration // 最多中毒时间
	for i := top; 0 < i; i-- {
		cntCur := timeSeries[i] - timeSeries[i-1] // 差值
		if cntCur < duration {
			cntMax -= duration
			cntMax += cntCur
		}
	}
	return cntMax
}
