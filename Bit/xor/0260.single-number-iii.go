package main

// https://leetcode.cn/problems/single-number-iii

// ❓ 只出现过一次的数字 III
// ⚠️ 2个数字出现1次

func singleNumber(nums []int) []int {
	numDiff := 0
	for _, num := range nums {
		// 两两抵消
		numDiff ^= num
	}
	// 101
	// 110
	// 011

	// -num = 取反+1，取最低有效位
	// 11110000
	// 00001111 + 1
	// 00010000

	numDiff = numDiff & -numDiff // TODO 两数 数位差 ,根据一位的不同 切分不同数字
	var num1, num2 int
	for _, num := range nums {
		if 0 < num&numDiff { // 最低有效位 为1
			num1 ^= num
		} else { //  最低有效位 为0
			num2 ^= num
		}
	}
	return []int{num1, num2}
}
