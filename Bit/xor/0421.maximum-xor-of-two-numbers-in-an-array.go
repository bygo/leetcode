package main

// https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array

// O( 30*n )

func findMaximumXOR(nums []int) int {
	var numRes, numMax int
	// 从高到低 按位确认
	for idx := 30; 0 <= idx; idx-- {
		// 结果翻倍
		numRes <<= 1
		// 下一个最大值
		numMax = numRes + 1

		// 每个数移位后 与 当前最大值 异或
		numMp := map[int]bool{}
		for _, num := range nums {
			// TODO 异或结合律 寻找另一个数 前缀
			num >>= idx
			numTarget := num ^ numMax
			if numMp[numTarget] {
				numRes += 1
				break
			}
			numMp[num] = true
		}
	}
	return numRes
}
