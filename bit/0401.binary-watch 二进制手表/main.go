package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/binary-watch

func readBinaryWatch(turnedOn int) []string {
	var res []string
	for h := 0; h < 12; h++ {
		for m := 0; m < 60; m++ {
			if OnesCounter(h)+OnesCounter(m) == turnedOn {
				res = append(res, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return res
}

func OnesCounter(i int) int {
	var ones int
	for 0 < i {
		ones++
		i = i & (i - 1)
	}
	return ones
}

func readBinaryWatch(turnedOn int) []string {
	var res []string
	for i := 0; i < 1024; i++ {
		h, m := i>>6, i&63 // 用位运算取出高 4 位和低 6 位
		if h < 12 && m < 60 && OnesCounter(i) == turnedOn {
			res = append(res, fmt.Sprintf("%d:%02d", h, m))
		}
	}
	return res
}
