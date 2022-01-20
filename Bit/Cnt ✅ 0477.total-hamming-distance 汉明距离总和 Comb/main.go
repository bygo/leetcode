package main

// https://leetcode-cn.com/problems/total-hamming-distance

// ❓ 汉明距离总和

func totalHammingDistance(nums []int) int {
	var cntTot int
	numsL := len(nums)

	// 枚举 1~30 位
	for pos := 0; pos < 30; pos++ {
		var cnt int
		// 统计1的个数
		for _, num := range nums {
			cnt += num >> pos & 1
		}
		// 乘以0的个数(组合)
		cntTot += cnt * (numsL - cnt)
	}
	return cntTot
}
