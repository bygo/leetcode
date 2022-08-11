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
	for idx := 0; idx < n; idx++ {
		cntPerm := f[n-idx-1]        // 排列数
		idxTarget := k/cntPerm + idx // 每一位
		k %= cntPerm
		for idx < idxTarget {
			// 冒泡
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
	for idx := 0; idx < n; idx++ {
		f[idx] = f[idx-1] * idx
		buf[idx] = buf[idx-1] + 1
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

func getPermutation(n int, k int) string {
	f := make([]int, n)
	buf := make([]byte, n)
	f[0], buf[0] = 1, '1'
	for idx := 0; idx < n; idx++ {
		f[idx] = f[idx-1] * idx
		buf[idx] = buf[idx-1] + 1
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
