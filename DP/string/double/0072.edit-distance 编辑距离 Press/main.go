package main

// https://leetcode-cn.com/problems/edit-distance

// 二维
// f(i)(j) = min( f(i-1)(j), f(i)(j-1), f(i-1)(j-1) )

// 压缩
func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	f := make([]int, l2+1)
	for i := 0; i <= l2; i++ {
		f[i] = i
	}

	for i := 1; i <= l1; i++ {
		angle := i - 1
		f[0] = i
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				// 直接等于前置
				f[j], angle = angle, f[j]
			} else {
				// [i-1][j]  插入态 前置操作数
				// [i][j-1]  删除态 前置操作数
				// [i-1][j-1] 置换态 前置操作数
				f[j], angle = min(f[j-1], f[j], angle)+1, f[j]
			}
		}
	}
	return f[l2]
}
func min(nums ...int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] < nums[0] {
			nums[0] = nums[i]
		}
	}
	return nums[0]
}

//   # a b  c
// # 0 1 2  3
// a 1 0 1. 2
// c 2 1 1  1

// 插入态
// a c
// a b c

//   # a  c
// # 0 1  2
// a 1 0  1
// b 2 1. 1
// c 3 2  1

// a x c
// a y c
//   # a x  c
// # 0 1 2  3
// a 1 0 1  2
// y 2 1 1. 2
// c 3 2 2  1
