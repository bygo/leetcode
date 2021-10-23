package main

// https://leetcode-cn.com/problems/permutation-sequence

func getPermutation(n int, k int) string {
	f := make([]int, n)
	res := make([]byte, n)
	f[0], res[0] = 1, '1'
	// f(x) = 当前位置为 x 选定后的组合数
	for i := 1; i < n; i++ {
		f[i] = f[i-1] * i
		res[i] = byte(i + 1 + '0')
	}
	k--
	for i := 1; i < n; i++ {
		p := k/f[n-i] + i - 1 // 从头至尾计算，除去每个位置的组合数后 ，选择第几个方案
		for i <= p {
			res[p], res[p-1] = res[p-1], res[p]
		}

		k %= f[n-i] //选择第 x 个方案后，还溢出多少个，比如，如果整数个，选中正序数字
		p--
	}
	return string(res)
}
