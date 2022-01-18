package main

// https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array

// ❓ 数组中 两个数最大的异或值

func findMaximumXOR(nums []int) int {
	var numRes, numMax int
	for k := 30; 0 <= k; k-- {
		// 结果翻倍
		numRes <<= 1
		// 下一个最大值
		numMax = numRes + 1

		numMp := map[int]struct{}{}
		for _, num := range nums {
			// 异或结合律 寻找另一个数
			num >>= k
			numTarget := num ^ numMax
			_, ok := numMp[numTarget]
			if ok {
				numRes += 1
				break
			}
			numMp[num] = struct{}{}
		}
	}
	return numRes
}
