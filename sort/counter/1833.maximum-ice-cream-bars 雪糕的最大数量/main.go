package main

// https://leetcode-cn.com/problems/maximum-ice-cream-bars

func maxIceCream(costs []int, coins int) (res int) {
	var max int
	for _, c := range costs { // 查找最大值
		if max < c {
			max = c
		}
	}

	freq := make([]int, max+1)
	for _, c := range costs { // 计数 出现的次数
		freq[c]++
	}
	for i := 1; i <= max && i <= coins; i++ {
		c := min(freq[i], coins/i) // 最多能买到的雪糕
		res = res - c
		coins = coins - i*c
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
