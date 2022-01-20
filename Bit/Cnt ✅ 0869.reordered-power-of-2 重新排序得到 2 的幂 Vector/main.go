package main

// https://leetcode-cn.com/problems/reordered-power-of-2

// ❓ 重新排序得到 2 的幂

func countDigits(n int) [10]int {
	// 统计总数
	var numMpCnt [10]int
	for 0 < n {
		numMpCnt[n%10]++
		n /= 10
	}
	return numMpCnt
}

var vecMpBool = map[[10]int]bool{}

func init() {
	for n := 1; n <= 1e9; n <<= 1 {
		// 2的幂 每数位的数字总数 vector
		vec := countDigits(n)
		vecMpBool[vec] = true
	}
}

func reorderedPowerOf2(n int) bool {
	vec := countDigits(n)
	return vecMpBool[vec]
}
