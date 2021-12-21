package main

// https://leetcode-cn.com/problems/maximum-ice-cream-bars

// ❓ 雪糕的最大数量

func maxIceCream(costs []int, coins int) int {
	// 查找最大值
	var costMax int
	for _, cost := range costs {
		if costMax < cost {
			costMax = cost
		}
	}

	// 计数 出现的次数
	var costMpCnt = make([]int, costMax+1)
	for _, cost := range costs {
		costMpCnt[cost]++
	}

	var cntMax int
	for coin := 1; coin <= costMax && coin <= coins; coin++ {
		cntMin := min(costMpCnt[coin], coins/coin) // 最多能买到的雪糕
		cntMax += cntMin
		coins -= coin * cntMin
	}
	return cntMax
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
