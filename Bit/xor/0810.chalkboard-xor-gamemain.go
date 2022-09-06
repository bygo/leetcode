package main

// https://leetcode.cn/problems/chalkboard-xor-game

func xorGame(nums []int) bool {
	if len(nums)%2 == 0 {
		// TODO 偶数长度，必定可以挑出一个数，使之不为0
		// 0 0 0 0  即为0 直接赢
		// 0 0 0 1 = 1
		// 0 1 1 1 = 1
		return true
	}
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	return xor == 0 // TODO 奇数长度，整体必须为0，否则当擦掉一个数，为偶数时，对方必胜
}
