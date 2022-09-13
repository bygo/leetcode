package main

import "strconv"

// https://leetcode.cn/problems/maximum-swap

func maximumSwap(num int) int {
	buf := []byte(strconv.Itoa(num))
	bufL := len(buf)
	idxMax, idxLo, idxHi := bufL-1, -1, -1
	for idx := bufL - 1; 0 <= idx; idx-- {
		if buf[idxMax] < buf[idx] {
			idxMax = idx
		} else if buf[idx] < buf[idxMax] { // TODO 贪心，大数位有更小的数
			idxLo, idxHi = idx, idxMax
		}
	}
	if idxLo == -1 {
		return num
	}

	buf[idxLo], buf[idxHi] = buf[idxHi], buf[idxLo]
	num, _ = strconv.Atoi(string(buf))
	return num
}
