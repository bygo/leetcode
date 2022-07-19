package main

// https://leetcode.cn/problems/coin-change-2

// 2 1   1 2
// 先放入1 再放入2 再放入3

// f[i] = f[i-coin] + 1
func change(amount int, coins []int) int {
	f := make([]int, amount+1)
	f[0] = 1

	// 多加一种硬币 ，f[i] 可以多几种跳跃方式
	// 1 1 1
	// 1 2
	// 2 1 ❌

	// 1 1 1 1
	// 1 1 2
	// 1 2 1 ❌
	// 2 1 1 ❌
	// 2 2

	for _, coin := range coins {
		for i := 1; i <= amount; i++ {
			if coin <= i {
				f[i] += f[i-coin]
			}
		}
	}
	return f[amount]
}

func wrong(amount int, coins []int) int {
	f := make([]int, amount+1)
	f[0] = 1

	// 1 1 1
	// 1 2
	// 2 1 ✅

	// 1 1 1 1
	// 1 1 2
	// 1 2 1 ✅
	// 2 1 1 ✅
	// 2 2
	// 多加一块钱 f[i] 可以多几种跳跃方式  （斐波那契数）
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				f[i] += f[i-coin]
			}
		}
	}
	return f[amount]
}
