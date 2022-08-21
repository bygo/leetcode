package main

// https://leetcode.cn/problems/permutation-sequence

func getPermutation(n int, k int) string {
	f := make([]int, n)
	buf := make([]byte, n)
	f[0], buf[0] = 1, '1'
	for idx := 1; idx < n; idx++ {
		f[idx] = f[idx-1] * idx
		buf[idx] = buf[idx-1] + 1
	}
	k--

	for idx := 0; idx < n-1; idx++ {
		cntPerm := f[n-idx-1] // Permutation the count
		idxTarget := k/cntPerm + idx
		k %= cntPerm
		for idx < idxTarget {
			// Bubbling
			buf[idxTarget], buf[idxTarget-1] = buf[idxTarget-1], buf[idxTarget]
			idxTarget--
		}
	}
	return string(buf)
}
