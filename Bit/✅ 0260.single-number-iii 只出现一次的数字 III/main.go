package main

// https://leetcode-cn.com/problems/single-number-iii

// ❓ 只出现过一次的数字 III
// ⚠️ 2个数字出现1次

func singleNumber(nums []int) []int {
	num := 0
	for i := range nums {
		// 两两抵消
		num ^= nums[i]
	}

	// -num,取反+1，最低有效位
	// 1111000
	// 0000111 + 1
	num = num & -num // 数字不同 肯定会有一位不同
	var num1, num2 int
	for i := range nums {
		if 0 < nums[i]&num {
			num1 ^= nums[i]
		} else {
			num2 ^= nums[i]
		}
	}
	return []int{num1, num2}
}
