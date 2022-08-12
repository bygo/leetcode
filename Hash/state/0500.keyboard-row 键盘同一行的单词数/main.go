package main

// https://leetcode.cn/problems/keyboard-row

// ❓ 在键盘同一行的单词数

var chMpState = map[byte]uint8{}

func init() {
	for _, ch := range []byte("qwertyuiopQWERTYUIOP") {
		chMpState[ch] = 1
	}
	for _, ch := range []byte("asdfghjklASDFGHJKL") {
		chMpState[ch] = 2
	}
	for _, ch := range []byte("zxcvbnmZXCVBNM") {
		chMpState[ch] = 3
	}
}

func findWords(words []string) []string {
	var wordSame []string
	for _, word := range words {
		wordL := len(word)
		chZero := word[0]
		idx := 1
		for idx < wordL {
			chCur := word[idx]
			if chMpState[chCur] != chMpState[chZero] {
				break
			}
			idx++
		}
		if idx == wordL {
			wordSame = append(wordSame, word)
		}
	}
	return wordSame
}
