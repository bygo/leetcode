package main

// https://leetcode-cn.com/problems/keyboard-row

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
		idx := 0
		wordL := len(word)
		for idx < wordL && chMpState[word[idx]] == chMpState[word[0]] {
			idx++
		}
		if idx == wordL {
			wordSame = append(wordSame, word)
		}
	}
	return wordSame
}
