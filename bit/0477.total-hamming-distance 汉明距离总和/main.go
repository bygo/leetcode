package main

// https://leetcode-cn.com/problems/total-hamming-distance

func totalHammingDistance(nums []int) (ans int) {
	n := len(nums)
	for i := 0; i < 30; i++ {
		c := 0
		for _, val := range nums {
			c += val >> i & 1 // 逐位计算 有多少个1
		}
		ans += c * (n - c) // 0*1
	}
	return
}
