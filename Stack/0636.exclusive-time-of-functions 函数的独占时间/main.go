package main

import (
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/exclusive-time-of-functions

func exclusiveTime(n int, logs []string) []int {
	res := make([]int, n)
	space := -1
	var stack []int
	for i := range logs {
		log := strings.Split(logs[i], ":")
		time, _ := strconv.Atoi(log[2])

		top := len(stack) - 1
		if log[1] == "start" {
			if -1 < top {
				res[stack[top]] += time - space // 前面一个开始了多久
			}
			index, _ := strconv.Atoi(log[0])
			stack = append(stack, index)
			space = time
		} else if log[1] == "end" {
			index := stack[top]
			stack = stack[:top]
			res[index] += time + 1 - space // 运行了多久
			space = time + 1
		}
	}
	return res
}
