package main

// https://leetcode-cn.com/problems/counting-bits

// ❓ 比特位计数
// ⚠️ 0~n 有几个1

func countBits(numMax int) []int {
	var bits []int
	for num := 0; num <= numMax; num++ {
		bits = append(bits, bk(num))
	}
	return bits
}

func bk(num int) int {
	var cnt int
	for 0 < num {
		num &= num - 1
		cnt++
	}
	return cnt
}

// 最高有效位
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

// 最低有效位
func countBits(numMax int) []int {
	var bits = make([]int, numMax+1)
	for num := 1; num <= numMax; num++ {
		bits[num] = bits[num>>1] + num&1
	}
	return bits
}

// 最低设置位
func countBits(numMax int) []int {
	var bits = make([]int, numMax+1)
	for num := 1; num <= numMax; num++ {
		bits[num] = bits[num&(num-1)] + 1
	}
	return bits
}
