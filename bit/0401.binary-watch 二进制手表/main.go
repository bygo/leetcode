package main

import (
	"fmt"
	"math/bits"
)

// Link: https://leetcode-cn.com/problems/binary-watch

func readBinaryWatch(turnedOn int) []string {
	var res []string
	for h := uint8(0); h < 12; h++ {
		for m := uint8(0); m < 60; m++ {
			if bits.OnesCount8(h)+bits.OnesCount8(m) == turnedOn {
				res = append(res, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return res
}

func readBinaryWatch(turnedOn int) []string {
	var res []string
	for i := 0; i < 1024; i++ {
		h, m := i>>6, i&63 // 用位运算取出高 4 位和低 6 位
		if h < 12 && m < 60 && bits.OnesCount(uint(i)) == turnedOn {
			res = append(res, fmt.Sprintf("%d:%02d", h, m))
		}
	}
	return res
}
