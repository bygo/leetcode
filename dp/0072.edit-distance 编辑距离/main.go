package main

//Link: https://leetcode-cn.com/problems/edit-distance

// dp
func minDistanc(word1, word2 string) int {
	m, n := len(word1), len(word2)
	var dp = make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for i := range dp[0] {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			angle := dp[i-1][j-1]
			if word1[i-1] == word2[j-1] {
				angle -= 1
			}
			dp[i][j] = min(dp[i-1][j], dp[i][j-1], angle) + 1
		}
	}

	return dp[m][n]
}

// 压缩
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	var dp = []int{}
	for i := 0; i <= n; i++ {
		dp = append(dp, i)
	}

	for i := 1; i <= m; i++ {
		angle := i - 1
		dp[0] = i
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				angle -= 1
			}
			dp[j], angle = min(dp[j-1], dp[j], angle)+1, dp[j]
		}
	}

	// dp[i-1][j-1] 不同等于替换
	// dp[i][j-1] 插入2
	// dp[i-1][j] 插入1
	// 	  #  h  h  h  o
	// #  0  1  2  3  4
	// h  1  0  1  2  3
	// o  2  1  1  2  2

	//    #   s   i   t   t   e   m
	// #  0   1   2   3   4   5   6
	// k  1   1   2   3   4   5   6
	// i  2   2   1   2   3   4   5
	// t  3   3
	// t  4
	// e  5
	// m  6
	return dp[n]
}

func min(nums ...int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] < nums[0] {
			nums[0] = nums[i]
		}
	}
	return nums[0]
}
