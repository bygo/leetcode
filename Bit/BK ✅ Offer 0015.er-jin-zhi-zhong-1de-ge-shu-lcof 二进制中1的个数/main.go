package main

// https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/submissions/

// ❓ 二进制中1 的个数

func hammingWeight(num uint32) int {
	var ones int
	for 0 < num {
		ones++
		// 每次消去最右的1
		num &= num - 1
	}
	return ones
}
