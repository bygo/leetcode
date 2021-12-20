package main

// https://leetcode-cn.com/problems/find-the-town-judge

func findJudge(n int, trust [][]int) int {
	out := map[int]int{}
	in := map[int]int{}
	for i := range trust {
		out[trust[i][0]]++ // 出度
		in[trust[i][1]]++  // 入度
	}

	for i := 1; i <= n; i++ {
		if out[i] == 0 && in[i] == n-1 {
			return i
		}
	}

	return -1
}
