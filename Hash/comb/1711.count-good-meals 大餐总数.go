package main

// https://leetcode.cn/problems/count-good-meals

// ❓ 大餐(和为2的幂)总数

func countPairs(deliciousness []int) int {
	var cnt, sumMax int
	numMpCnt := map[int]int{}
	for _, delicious := range deliciousness {
		// 最大幂
		if sumMax>>1 < delicious {
			sumMax = delicious << 1
		}
		// 2的幂循环一遍
		for sum := 1; sum <= sumMax; sum <<= 1 {
			if sum < delicious {
				continue
			}
			numTarget := sum - delicious
			cnt += numMpCnt[numTarget]
		}
		// 存入
		numMpCnt[delicious]++
	}
	return cnt % (1e9 + 7)
}
