package main

import (
	"fmt"
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
	return fmt.Sprintf("%d.%d.%d.%d", x>>24&0xff, x>>16&0xff, x>>8&0xff, x&0xff)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func lenBit(x int) int {
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
	var strs []string
	start := ipToInt(ip)
	if start == 0 {
		mask := 32 - lenBit(n) // TODO
		strs = append(strs, intToIP(start)+"/"+strconv.Itoa(mask))
		step := 1 << (32 - mask)
		start += step
		n -= step
	}
	for 0 < n {
		mask := 32 - min(lenBit(start&-start), lenBit(n)) // TODO
		strs = append(strs, intToIP(start)+"/"+strconv.Itoa(mask))
		step := 1 << (32 - mask)
		start += step
		n -= step
	}
	return strs
}
