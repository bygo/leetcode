package main

// https://leetcode-cn.com/problems/number-of-1-bits/
// ❓ 位1的个数

func hammingWeight(num uint32) int {
	cnt := 0
	for 0 < num {
		num = num & (num - 1)
		cnt++
	}
	return cnt
}
