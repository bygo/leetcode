package main

// https://leetcode.cn/problems/valid-anagram

// ❓ 有效的异位词

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	chMpCnt := [26]int{}
	for i := range s {
		chMpCnt[s[i]-'a']++
		chMpCnt[t[i]-'a']--
	}

	for _, cnt := range chMpCnt {
		if cnt != 0 {
			return false
		}
	}
	return true
}

// ❌ 通过字母 code 总和 求差，比如 bc 与 ad 相等， 却不是异位词
