package main

// https://leetcode.cn/problems/beautiful-arrangement

func countArrangement(n int) int {
	var cnt int
	var used int
	match := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i%j == 0 || j%i == 0 {
				match[i] = append(match[i], j) // TODO 预处理
			}
		}
	}
	var dfs func(idx int)
	dfs = func(idx int) {
		if n < idx {
			cnt++
			return
		}

		for _, num := range match[idx] {
			if used>>num&1 == 1 { // TODO 已被选择
				continue
			}
			used |= 1 << num // TODO
			dfs(idx + 1)
			used ^= 1 << num // TODO
		}
	}
	dfs(1)
	return cnt
}
