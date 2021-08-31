package main

// https://leetcode-cn.com/problems/total-hamming-distance

func totalHammingDistance(nums []int) int {
	var res int
	l1 := len(nums)
	for i := 0; i < 30; i++ {
		var cnt int
		for _, num := range nums {
			cnt += num >> i & 1
		}
		res += cnt * (l1 - cnt)
	}
	return res
}
