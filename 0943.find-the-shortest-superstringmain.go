package main

import "math"

// https://leetcode.cn/problems/find-the-shortest-superstring

func shortestSuperstring(words []string) string {
	max := math.MaxInt32
	wL := len(words)
	numMax := 1 << wL
	graph := make([][]int, wL)
	for i := 0; i < wL; i++ {
		graph[i] = make([]int, wL)
		for j := 0; j < wL; j++ {
			if i == j {
				continue
			}
			graph[i][j] = overlap(words[i], words[j]) // TODO i后缀 和 j前缀 重叠长度
		}
	}

	// TODO 数位dp 初始化
	f := make([][]int, numMax)
	//lasts := make([][]int, numMax)
	for i := 0; i < numMax; i++ {
		f[i] = make([]int, wL)
		//lasts[i] = make([]int, wL)
		for j := 0; j < wL; j++ {
			f[i][j] = max
			//lasts[i][j] = -1
		}
	}

	// 一位 状态
	for i := 0; i < wL; i++ {
		f[1<<i][i] = len(words[i])
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
				curL := f[subset][i] + len(words[j]) - graph[i][j]
				subsetNext := subset | 1<<j
				if curL < f[subsetNext][j] {
					f[subsetNext][j] = curL
					//lasts[subsetNext][j] = i
				}
			}
		}
	}
	// 最后一位
	valMax := math.MaxInt32
	var idxLast int
	for idx, val := range f[numMax-1] {
		if val < valMax {
			valMax = val
			idxLast = idx
		}
	}

	idxes := make([]int, wL)
	idxes[wL-1] = idxLast
	sub := numMax - 1
	for i := wL - 2; 0 <= i; i-- {
		for j, length := range f[sub] {
			tar := f[sub][idxLast] - len(words[j]) + graph[j][idxLast]
			if length == tar {
				idxLast = j
				break
			}
		}
		sub ^= 1 << idxLast
		//tmp := idxLast
		//idxLast = lasts[sub][idxLast]
		idxes[i] = idxLast
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
	var cnt int
	for idx := min(aL, bL) - 1; 0 <= idx; idx-- {
		if a[aL-idx:] == b[:idx] {
			cnt = idx
			break
		}
	}
	return cnt
}
