package main

// https://leetcode-cn.com/problems/count-good-meals

// ❓ 大餐(和为2的幂)总数

func countPairs(deliciousness []int) int {
	var cnt, sumMax int
	numMpCnt := map[int]int{}
	for _, delicious := range deliciousness {
		// 最大幂
		if sumMax>>1 < delicious {
			sumMax = delicious << 1
		}
		// 2的幂循环一遍s
		for sum := 1; sum <= sumMax; sum <<= 1 {
			numTarget := sum - delicious
			cnt += numMpCnt[numTarget]
		}
		// 存入
		numMpCnt[delicious] += 1
	}
	return cnt % (1e9 + 7)
}
