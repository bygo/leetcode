package main

// https://leetcode-cn.com/problems/single-number-ii

// ❓ 只出现一次的数字 II
// ⚠️ 其他出现3次

func singleNumber_(nums []int) int {
	a, b := 0, 0
	for _, num := range nums {
		b = (b ^ num) &^ a
		a = (a ^ num) &^ b
	}
	return b
}

// 依次确认

func singleNumber(nums []int) int {
	single := int32(0)
	for i := 0; i < 32; i++ {
		// 每位次数
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}
		// 1 4 7 ...
		if 0 < total%3 {
			single |= 1 << i
		}
	}
	return int(single)
}
