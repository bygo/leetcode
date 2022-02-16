package main

// https://leetcode-cn.com/problems/unique-binary-search-trees/

// ❓ 不同的二叉搜索树

func numTrees(n int) int {
	catalan := 1
	for i := 0; i < n; i++ {
		catalan = catalan * 2 * (2*i + 1) / (i + 2)
	}
	return catalan
}

func numTrees_DP(n int) int {
	f := make([]int, n+1)
	f[0], f[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			f[i] += f[j-1] * f[i-j]
		}
	}
	return f[n]
}
