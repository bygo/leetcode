package main

// https://leetcode-cn.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers

// ❓ 十-二进制数的最少数目
// 123 => 111 011 001 => 3

func minPartitions(str string) int {
	var max rune = -1
	for _, ch := range str {
		if max < ch {
			max = ch
		}
	}
	return int(max - '0')
}
