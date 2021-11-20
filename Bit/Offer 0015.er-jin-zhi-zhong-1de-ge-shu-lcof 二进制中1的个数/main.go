package main

// https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/submissions/

func hammingWeight(num uint32) int {
	var ones int
	for 0 < num {
		ones++
		num &= num - 1
	}
	return ones
}
