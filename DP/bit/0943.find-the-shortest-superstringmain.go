package bit

import (
	"math"
)

// https://leetcode.cn/problems/find-the-shortest-superstring

func shortestSuperstring(words []string) string {
	max := math.MaxInt32
	wL := len(words)
	numMax := 1 << wL

	// TODO 初始化
	f := make([][]int, numMax)
	for i := 0; i < numMax; i++ {
		f[i] = make([]int, wL)
		for j := 0; j < wL; j++ {
			f[i][j] = max
		}
	}

	graph := make([][]int, wL)
	for i := 0; i < wL; i++ {
		f[1<<i][i] = len(words[i]) // TODO 初始化

		graph[i] = make([]int, wL)
		for j := 0; j < wL; j++ {
			if i == j {
				continue
			}
			graph[i][j] = overlap(words[i], words[j]) // TODO i后缀 和 j前缀 重叠长度
		}
	}

	for subset := 0; subset < numMax-1; subset++ {
		for i := 0; i < wL; i++ {
			if f[subset][i] == max {
				continue
			}
			for j := 0; j < wL; j++ {
				if subset>>j&1 == 1 {
					continue
				}
				f[subset|1<<j][j] = min(f[subset|1<<j][j], f[subset][i]+len(words[j])-graph[i][j])
			}
		}
	}
	sub := numMax - 1
	var idxLast int
	for i := 0; i < wL; i++ {
		if f[sub][i] < f[sub][idxLast] {
			idxLast = i
		}
	}

	var idxes = make([]int, wL)
	for i := 0; i < wL; i++ {
		for j := 0; j < wL; j++ {
			if f[sub][j]+len(words[idxLast])-graph[j][idxLast] == f[sub|1<<idxLast][idxLast] {
				idxLast = j
				break
			}
		}
		sub ^= 1 << idxLast
		idxes[wL-i-1] = idxLast
	}
	var buf []byte
	buf = append(buf, words[idxes[0]]...)
	for i := 1; i < wL; i++ {
		word := words[idxes[i]]
		idxOver := graph[idxes[i-1]][idxes[i]]
		buf = append(buf, word[idxOver:]...)
	}
	return string(buf)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func overlap(a, b string) int {
	aL := len(a)
	bL := len(b)
	for idx := min(aL, bL); 0 <= idx; idx-- {
		if a[aL-idx:] == b[:idx] {
			return idx
		}
	}
	return 0
}
