package main

// https://leetcode.cn/problems/permutation-sequence

func getPermutation(n int, k int) string {
	f := make([]int, n)
	buf := make([]byte, n)
	f[0], buf[0] = 1, '1'
	for idx := 1; idx < n; idx++ {
		f[idx] = f[idx-1] * idx   // TODO 排列数 1*2*3*4*5
		buf[idx] = buf[idx-1] + 1 // TODO 第一个排列
	}
	k--

	for idx := 0; idx < n-1; idx++ {
		cntPerm := f[n-idx-1]
		idxTarget := idx + k/cntPerm // TODO 每个数 能有N个排列
		k %= cntPerm
		for idx < idxTarget {
			// Bubbling
			buf[idxTarget], buf[idxTarget-1] = buf[idxTarget-1], buf[idxTarget]
			idxTarget--
		}
	}
	return string(buf)
}
