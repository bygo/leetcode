package main

// https://leetcode-cn.com/problems/count-good-meals

func countPairs(deliciousness []int) (res int) {
	var preMax, sumMax int
	cnt := map[int]int{}
	for _, delicious := range deliciousness {
		if preMax < delicious {
			preMax = delicious
			sumMax = preMax << 1
		}
		for sum := 1; sum <= sumMax; sum <<= 1 {
			res += cnt[sum-delicious]
		}
		cnt[delicious] += 1
	}
	return res % (1e9 + 7)
}
