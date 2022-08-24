package main

// https://leetcode.cn/problems/number-of-1-bits/

func hammingWeight(num uint32) int {
	cnt := 0
	for 0 < num {
		num &= num - 1
		cnt++
	}
	return cnt
}
