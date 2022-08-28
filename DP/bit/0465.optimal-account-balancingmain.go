package bit

import (
	"math"
	"math/bits"
)

// https://leetcode.cn/problems/optimal-account-balancing

func minTransfers(transactions [][]int) int {
	cnts := []int{}
	for _, p := range transactions {
		for len(cnts) <= p[0] {
			cnts = append(cnts, 0)
		}
		for len(cnts) <= p[1] {
			cnts = append(cnts, 0)
		}
		cnts[p[0]] -= p[2]
		cnts[p[1]] += p[2]
	}

	cL := len(cnts)
	subMax := 1 << cL
	f := make([]int, subMax)
	for subset := 1; subset < subMax; subset++ {
		sum := 0
		for idx, cnt := range cnts {
			sum += subset >> idx & 1 * cnt
		}
		if sum != 0 {
			f[subset] = math.MaxInt32
			continue
		}
		f[subset] = bits.OnesCount(uint(subset)) - 1
		sub := subset
		for sub != 0 {
			sub = (sub - 1) & subset
			f[subset] = min(f[subset], f[sub]+f[subset^sub]) // TODO
		}
	}
	return f[subMax-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
