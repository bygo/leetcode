package main

// https://leetcode.cn/problems/reordered-power-of-2

func countDigits(numDec int) [10]int {
	// 统计总数
	var numMpCnt [10]int
	for 0 < numDec {
		numMpCnt[numDec%10]++
		numDec /= 10
	}
	return numMpCnt
}

var vecMpBool = map[[10]int]bool{}

func init() {
	for num := 1; num <= 1e9; num <<= 1 {
		// 2的幂 每数位的数字总数 vector
		vec := countDigits(num)
		vecMpBool[vec] = true
	}
}

func reorderedPowerOf2(n int) bool {
	vec := countDigits(n)
	return vecMpBool[vec]
}
