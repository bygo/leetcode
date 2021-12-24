package main

// https://leetcode-cn.com/problems/repeated-dna-sequences

// ❓ 统计重复的DNA 序

var bases = map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}

const baseL = 10
const base = 4

func findRepeatedDnaSequences(s string) []string {
	sL := len(s)
	if sL <= baseL {
		return nil
	}
	mul := pow(4, baseL-1) // 最高位
	hashCur := 0
	hashMpCnt := map[int]int8{}
	for i := 0; i < baseL; i++ {
		ch := s[i]
		bitLow := bases[ch]
		hashCur = hashCur*base + bitLow
	}
	hashMpCnt[hashCur]++

	// 统计
	var strsRepeat []string
	for idx := 0; idx < sL-baseL; idx++ {
		bitHigh := bases[s[idx]]      // 高位
		bitLow := bases[s[idx+baseL]] // 低位

		hashCur = hashCur - bitHigh*mul // 减最左
		hashCur *= base                 // 提升
		hashCur += bitLow               // 加最右

		// 临界
		if hashMpCnt[hashCur] < 2 {
			hashMpCnt[hashCur]++
			if hashMpCnt[hashCur] == 2 {
				strsRepeat = append(strsRepeat, s[idx+1:idx+baseL+1]) // 下一个字符串
			}
		}
	}
	return strsRepeat
}

func pow(x, n uint) int {
	var num uint = 1
	for 0 < n {
		if n&1 == 1 {
			num *= x
		}
		n >>= 1
		x *= x
	}
	return int(num)
}
