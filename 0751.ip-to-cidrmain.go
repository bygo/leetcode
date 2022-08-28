package main

import (
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/ip-to-cidr

func ipToInt(ip string) int {
	total := 0
	for _, block := range strings.Split(ip, ".") {
		num, _ := strconv.Atoi(block)
		total = total<<8 + num
	}
	return total
}
func intToIP(x int) string {
	ip := ""
	for i := 3; 0 <= i; i-- {
		ip += strconv.Itoa((x>>(i*8))&0xff) + "."
	}
	return ip[:len(ip)-1]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func length(x int) int {
	lo, hi := -1, 32
	for lo < hi {
		mid := int(uint(lo+hi+1) >> 1)
		if x>>mid == 0 {
			hi = mid - 1
		} else {
			lo = mid
		}
	}
	return lo
}

func ipToCIDR(ip string, n int) []string {
	start := ipToInt(ip)
	var strs []string
	for 0 < n {
		mask := max(32-length(start&-start), 32-length(n))
		if start == 0 {
			mask = 32 - length(n)
		}
		strs = append(strs, intToIP(start)+"/"+strconv.Itoa(mask))
		step := 1 << (32 - mask)
		start += step
		n -= step
	}
	return strs
}
