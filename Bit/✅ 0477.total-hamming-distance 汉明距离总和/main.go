package main

// https://leetcode-cn.com/problems/total-hamming-distance

// ❓ 汉明距离总和

func totalHammingDistance(nums []int) int {
	var cntTot int
	numsL := len(nums)
	for i := 0; i < 30; i++ {
		var cnt int
		// 统计1的个数
		for _, num := range nums {
			cnt += num >> i & 1
		}
		// 乘以0的个数(组合)
		cntTot += cnt * (numsL - cnt)
	}
	return cntTot
}
