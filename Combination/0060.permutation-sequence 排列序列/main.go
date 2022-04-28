package main

// https://leetcode-cn.com/problems/permutation-sequence

func getPermutation(n int, k int) string {
	f := make([]int, n)
	buf := make([]byte, n)
	f[0], buf[0] = 1, '1'
	for idx := 1; idx < n; idx++ {
		f[idx] = f[idx-1] * idx
		buf[idx] = byte(idx + '1')
	}
	k--
	for idx := 0; idx < n; idx++ {
		cntPerm := f[n-idx-1]
		idxTarget := k/cntPerm + idx
		k %= cntPerm
		for idx < idxTarget {
			buf[idxTarget], buf[idxTarget-1] = buf[idxTarget-1], buf[idxTarget]
			idxTarget--
		}
	}
	return string(buf)
}
