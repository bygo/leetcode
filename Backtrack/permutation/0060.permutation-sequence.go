package main

// https://leetcode.cn/problems/permutation-sequence

func getPermutation(n int, k int) string {
	f := make([]int, n)
	buf := make([]byte, n)
	f[0], buf[0] = 1, '1'
	// 1~5
	for idx := 1; idx < n; idx++ {
		f[idx] = f[idx-1] * idx   // TODO 1*1*2*3*4
		buf[idx] = buf[idx-1] + 1 // TODO 第1个排列
	}
	k--

	for idx := 0; idx < n-1; idx++ {
		cntPerm := f[n-idx-1]
		idxTarget := idx + k/cntPerm // TODO 每个数 能包含N个排列
		k %= cntPerm                 // TODO 剩下的排列数
		for idx < idxTarget {
			// Bubbling
			buf[idxTarget], buf[idxTarget-1] = buf[idxTarget-1], buf[idxTarget]
			idxTarget--
		}
	}
	return string(buf)
}

func getPermutation(n int, k int) string {
	f := make([]int, n)
	buf := make([]byte, n)
	f[0], buf[0] = 1, '1'
	// 1~5
	for idx := 1; idx < n; idx++ {
		f[idx] = f[idx-1] * idx
		buf[idx] = buf[idx-1] + 1
	}
	k--

	var bufRes = []byte{}
	for idx := 0; idx < n && 0 < k; idx++ {
		cntPerm := f[n-idx-1]
		target := k / cntPerm
		bufRes = append(bufRes, buf[target])
		buf = append(buf[:target], buf[target+1:]...)
		k %= cntPerm
	}
	bufRes = append(bufRes, buf...)
	return string(bufRes)
}
