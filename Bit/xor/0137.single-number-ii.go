package main

// https://leetcode.cn/problems/single-number-ii

// ❓ 只出现一次的数字 II
// ⚠️ 其他出现3次
// 		lo:		       hi:
// 1: 0^1 &^ 0 = 1 | 0^1 &^ 1 = 0 // lo ^进位, hi &^抵消
// 0: 1^0 &^ 0 = 1 | 0^0 &^ 1 = 0
// 1: 1^1 &^ 0 = 0 | 0^1 &^ 0 = 1 // lo ^抵消，hi ^进位
// 0: 0^0 &^ 1 = 0 | 1^0 &^ 0 = 1
// 1: 0^1 &^ 1 = 0 | 1^1 &^ 0 = 0 // lo &^抵消,hi ^抵消
// 0: 0^0 &^ 0 = 0 | 0^0 &^ 0 = 0

func singleNumber(nums []int) int {
	var lo, hi int
	for _, num := range nums {
		lo = (lo ^ num) &^ hi // TODO 第三次lo 被抑制 为0
		hi = (hi ^ num) &^ lo // 第一次被 lo 抑制 为0，第二次没被 lo 抑制 为1，
	}
	return lo
}

// 依位确认

func singleNumber(nums []int) int {
	single := int32(0)
	for pos := 0; pos < 32; pos++ {
		// 每位次数
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> pos & 1
		}
		// 1 4 7 ...
		if 0 < total%3 { // TODO 分治
			single |= 1 << pos
		}
	}
	return int(single)
}
