package main

// https://leetcode.cn/problems/edit-distance

// 二维
// f(i)(j) = min( f(i-1)(j), f(i)(j-1), f(i-1)(j-1) )
func minDistance(word1, word2 string) int {
	l1, l2 := len(word1), len(word2)
	var f = make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		f[i] = make([]int, l2+1)
		// 空字符串操作数
		f[i][0] = i
	}

	// 空字符串 操作数
	for i := range f[0] {
		f[0][i] = i
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				// 相等时,直接等于前置
				f[i][j] = f[i-1][j-1]
			} else {
				// 删除1 [i-1][j]的距离
				// 删除2 [i][j-1]的距离
				// 替换 [i-1][j-1]的距离
				f[i][j] = min(f[i-1][j], f[i][j-1], f[i-1][j-1]) + 1
			}
		}
	}
	return f[l1][l2]
}

func min(nums ...int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] < nums[0] {
			nums[0] = nums[i]
		}
	}
	return nums[0]
}
