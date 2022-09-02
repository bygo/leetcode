package main

import (
	"fmt"
)

// https://leetcode.cn/problems/binary-watch

func readBinaryWatch(turnedOn int) []string {
	var strsTime []string
	for hour := 0; hour < 12; hour++ {
		for min := 0; min < 60; min++ {
			if OnesCounter(hour)+OnesCounter(min) == turnedOn {
				strsTime = append(strsTime, fmt.Sprintf("%d:%02d", hour, min))
			}
		}
	}
	return strsTime
}

func OnesCounter(num int) int {
	var ones int
	for 0 < num {
		ones++
		num &= num - 1
	}
	return ones
}

func readBinaryWatch(turnedOn int) []string {
	var strsTime []string
	for num := 0; num < 1<<11; num++ {
		hour, min := num>>6, num&0b111111
		if hour < 12 && min < 60 && OnesCounter(num) == turnedOn {
			strsTime = append(strsTime, fmt.Sprintf("%d:%02d", hour, min))
		}
	}
	return strsTime
}
