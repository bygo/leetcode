package main

// https://leetcode.cn/problems/day-of-the-year

var monthMpDays = [12]int{
	31,
	28,
	31,
	30,
	31,
	30,
	31,
	31,
	30,
	31,
	30,
	31,
}

func dayOfYear(date string) int {
	var year, month, day int
	var idx int

	for date[idx] != '-' {
		year = year*10 + int(date[idx]-'0')
		idx++
	}

	idx++

	for date[idx] != '-' {
		month = month*10 + int(date[idx]-'0')
		idx++
	}
	idx++

	for idx < len(date) {
		day = day*10 + int(date[idx]-'0')
		idx++
	}

	var cntDay int
	for monthIdx := 0; monthIdx < month-1; monthIdx++ {
		cntDay += monthMpDays[monthIdx]
	}

	cntDay += day

	if 2 < month && year%4 == 0 {
		cntDay++
	}

	return cntDay
}
