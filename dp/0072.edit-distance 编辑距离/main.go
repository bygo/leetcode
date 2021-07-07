package main

// Link: https://leetcode-cn.com/problems/edit-distance

// f(i)(j) = min( f(i-1)(j), f(i)(j-1), f(i-1)(j-1) )
func minDistance(word1, word2 string) int {
	l1, l2 := len(word1), len(word2)
	var f = make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		f[i] = make([]int, l2+1)
		f[i][0] = i
	}

	for i := range f[0] {
		f[0][i] = i
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			angle := f[i-1][j-1]
			if word1[i-1] == word2[j-1] {
				angle -= 1
			}
			f[i][j] = min(f[i-1][j], f[i][j-1], angle) + 1
		}
	}

	return f[l1][l2]
}

// 压缩
func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	var f = []int{}
	for i := 0; i <= l2; i++ {
		f = append(f, i)
	}

	for i := 1; i <= l1; i++ {
		angle := i - 1
		f[0] = i
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				angle -= 1
			}
			f[j], angle = min(f[j-1], f[j], angle)+1, f[j]
		}
	}

	// f[i-1][j-1] 不同等于替换
	// f[i][j-1] 插入2
	// f[i-1][j] 插入1
	// 	  #  h  h  h  o
	// #  0  1  2  3  4
	// h  1  0  1  2  3
	// o  2  1  1  2  2

	//    #   s   i   t   t   e   l1
	// #  0   1   2   3   4   5   6
	// k  1   1   2   3   4   5   6
	// i  2   2   1   2   3   4   5
	// t  3   3
	// t  4
	// e  5
	// l1  6
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
