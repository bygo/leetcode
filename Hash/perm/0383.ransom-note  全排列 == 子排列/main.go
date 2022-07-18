package main

// https://leetcode.cn/problems/ransom-note

// ❓ 从magazine组合成ransomNote

func canConstruct(ransomNote string, magazine string) bool {
	chMpCnt := map[byte]int{}
	for i := range magazine {
		ch := magazine[i]
		chMpCnt[ch]++
	}

	for i := range ransomNote {
		ch := ransomNote[i]
		if chMpCnt[ch] == 0 {
			return false
		}
		chMpCnt[ch]--
	}
	return true
}
