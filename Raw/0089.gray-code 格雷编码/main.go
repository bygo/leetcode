package main

// https://leetcode-cn.com/problems/gray-code

// ❓ 格雷编码

func grayCode(n int) []int {
	f := []int{0}

	max := 1 << n
	for i := 1; i < max; i <<= 1 {
		for j := len(f) - 1; 0 <= j; j-- {
			f = append(f, f[j]|i)
		}
	}
	return f
}
