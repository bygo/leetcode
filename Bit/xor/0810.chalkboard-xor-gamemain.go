package main

// https://leetcode.cn/problems/chalkboard-xor-game

// 删除一个数值后为0 == 失败

func xorGame(nums []int) bool {
	if len(nums)%2 == 0 {
		// 0 0 0 1 = 1
		// 0 1 1 1 = 1
		// TODO 偶数长度，奇数1 必定挑出一个0，使之继续为奇数1，即不为0
		return true
	}
	xor := 0
	for _, num := range nums { // TODO 整体为0，擦掉任何数都将不为0
		xor ^= num
	}
	return xor == 0
}
