package main

// https://leetcode-cn.com/problems/reordered-power-of-2

// ❓ 重新排序得到 2 的幂

func countDigits(n int) (cnt [10]int) {
	// 统计总数
	for 0 < n {
		cnt[n%10]++
		n /= 10
	}
	return
}

var powerOf2Digits = map[[10]int]bool{}

// vector
func init() {
	for n := 1; n <= 1e9; n <<= 1 {
		// 2的幂 每数位的数字总数
		powerOf2Digits[countDigits(n)] = true
	}
}

func reorderedPowerOf2(n int) bool {
	return powerOf2Digits[countDigits(n)]
}
