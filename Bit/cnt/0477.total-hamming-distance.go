package main

// https://leetcode.cn/problems/total-hamming-distance

func totalHammingDistance(nums []int) int {
	var cntTot int
	numsL := len(nums)

	for idx := 0; idx < 30; idx++ {
		var cnt int
		for _, num := range nums {
			cnt += num >> idx & 1
		}
		cntTot += cnt * (numsL - cnt) // TODO
	}
	return cntTot
}
