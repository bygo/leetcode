package main

// https://leetcode-cn.com/problems/ransom-note

// ❓ 从magazine组合成ransomNote

func canConstruct(ransomNote string, magazine string) bool {
	chMpCnt := map[byte]int{}
	for i := range magazine {
		chMpCnt[magazine[i]]++
	}

	for i := range ransomNote {
		if chMpCnt[ransomNote[i]] == 0 {
			return false
		}
		chMpCnt[ransomNote[i]]--
	}
	return true
}
