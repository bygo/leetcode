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

	// 	  #  h  o  r
	// #  0  1  2  3
	// h  1  0  1  2
	// o  2  1  0  1

	//    #   s   i   t   t  e  n
	// #  0   1   2   3   4  5  6
	// k  1   1   2   3   4  5  6
	// i  2   2   1   2   3  4  5
	// t  3   3
	// t  4
	// e  5
	// n  6
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
