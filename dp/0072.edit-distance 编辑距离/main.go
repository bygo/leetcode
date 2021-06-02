package main

//Link: https://leetcode-cn.com/problems/edit-distance

func main() {
	println(minDistance("hor", "ho"))
}
func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	var dp = []int{}
	for i := 0; i <= m; i++ {
		dp = append(dp, i)
	}

	for i := 1; i <= n; i++ {
		angle := i - 1
		dp[0] = i
		for j := 1; j <= m; j++ {
			if word1[i-1] == word2[j-1] {
				angle -= 1
			}
			dp[j], angle = min(dp[j-1], dp[j], angle)+1, dp[j]
		}
	}

	return dp[m]
}

func min(nums ...int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] < nums[0] {
			nums[0] = nums[i]
		}
	}
	return nums[0]
}
