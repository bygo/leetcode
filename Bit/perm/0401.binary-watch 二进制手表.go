package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/binary-watch

// ❓ 二进制手表

// 枚举 h 0~12 m 0~60
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

// 枚举2^10
func readBinaryWatch_(turnedOn int) []string {
	var strsTime []string
	for num := 0; num < 1024; num++ {
		hour, min := num>>6, num&0b111111 // 用位运算取出高 4 位和低 6 位
		if hour < 12 && min < 60 && OnesCounter(num) == turnedOn {
			strsTime = append(strsTime, fmt.Sprintf("%d:%02d", hour, min))
		}
	}
	return strsTime
}
