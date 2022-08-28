package bit

// https://leetcode.cn/problems/counting-bits

// Highest significant bit
func countBits(numMax int) []int {
	var bits = make([]int, numMax+1)
	var bitHi = 0
	for num := 1; num <= numMax; num++ {
		if num&(num-1) == 0 {
			bitHi = num
		}
		bits[num] = bits[num-bitHi] + 1
	}
	return bits
}

// Lowest significant bit
func countBits(numMax int) []int {
	var bits = make([]int, numMax+1)
	for num := 1; num <= numMax; num++ {
		bits[num] = bits[num>>1] + num&1
	}
	return bits
}

// Lowest setting bit
func countBits(numMax int) []int {
	var bits = make([]int, numMax+1)
	for num := 1; num <= numMax; num++ {
		bits[num] = bits[num&(num-1)] + 1
	}
	return bits
}
