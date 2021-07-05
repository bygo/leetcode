package main

// Link: https://leetcode-cn.com/problems/pascals-triangle-ii

func getRow(rowIndex int) []int {
	numRows := rowIndex + 1
	var dp = make([]int, numRows)
	for i := 0; i < numRows; i++ {
		dp[0] = 1
		dp[i] = 1
		for j := i - 1; 0 < j; j-- {
			dp[j] = dp[j-1] + dp[j]
		}
	}
	return dp
}
