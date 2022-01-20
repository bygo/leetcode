package main

// https://leetcode-cn.com/problems/repeated-dna-sequences

// ❓ 统计重复的DNA 序

const baseL = 10
const bitStand = 1<<(baseL*2) - 1

//var base = map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}
var base = ['T' + 1]int32{
	'A': 0,
	'C': 1,
	'G': 2,
	'T': 3,
}

func findRepeatedDnaSequences(s string) []string {
	var strsRepeat []string
	sL := len(s)
	if sL <= baseL {
		return nil
	}
	var num int32
	for i := range s[:baseL-1] {
		num = num<<2 | base[s[i]]
	}

	numMpCnt := [1 << 20]int8{}
	for idx := 0; idx <= sL-baseL; idx++ {
		bitLo := s[idx+baseL-1]
		num = num<<2 | base[bitLo] // 提升
		num = num & bitStand       // 对齐
		if numMpCnt[num] < 2 {
			numMpCnt[num]++
			if numMpCnt[num] == 2 {
				strsRepeat = append(strsRepeat, s[idx:idx+baseL])
			}
		}

	}
	return strsRepeat
}
