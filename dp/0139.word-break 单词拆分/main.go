package main

// https://leetcode-cn.com/problems/word-break

// ❓ 单词拆分

func wordBreak(s string, wordDict []string) bool {
	wordMpBool := make(map[string]bool)
	for _, word := range wordDict {
		wordMpBool[word] = true
	}
	sL := len(s)
	idxMpBool := make([]bool, sL+1)
	idxMpBool[0] = true

	for r := 1; r <= sL; r++ {
		for l := 0; l < r; l++ {
			if idxMpBool[l] && wordMpBool[s[l:r]] {
				idxMpBool[r] = true
				break
			}
		}
	}

	return idxMpBool[sL]
}
