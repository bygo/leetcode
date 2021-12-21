package main

// https://leetcode-cn.com/problems/find-words-that-can-be-formed-by-characters

// ❓ 能拼写的所有单词长度之和

func countCharacters(words []string, chars string) int {
	chMpCnt := [26]int{}
	for i := range chars {
		ch := chars[i] - 'a'
		chMpCnt[ch]++
	}

	var cnt int
	for i := range words {
		chMpCntCur := [26]int{}
		word := words[i]
		for j := range word {
			ch := word[j] - 'a'
			chMpCntCur[ch]++
		}
		k := 0
		for k < 26 {
			if chMpCnt[k] < chMpCntCur[k] {
				break
			}
			k++
		}
		if k == 26 {
			cnt += len(words[i])
		}
	}
	return cnt
}
