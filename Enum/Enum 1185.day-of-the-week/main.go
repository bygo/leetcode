package main

// https://leetcode-cn.com/problems/day-of-the-week

var week = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
var monthDays = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30}

func dayOfTheWeek(day, month, year int) string {
	cntDay := 0
	// 年
	cntDay += 365*(year-1971) + (year-1969)/4 // 1973 年 算1972 多出一天

	// 月
	for _, d := range monthDays[:month-1] {
		cntDay += d
	}
	// 润
	if 3 <= month && (year%400 == 0 || year%4 == 0 && year%100 != 0) {
		cntDay++
	}
	// 日
	cntDay += day
	return week[(cntDay+3)%7]
}
