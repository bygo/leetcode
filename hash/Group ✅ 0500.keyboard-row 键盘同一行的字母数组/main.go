package main

// https://leetcode-cn.com/problems/keyboard-row

var m = map[byte]uint8{}

func init() {
	for _, v := range []byte("qwertyuiopQWERTYUIOP") {
		m[v] = 1
	}
	for _, v := range []byte("asdfghjklASDFGHJKL") {
		m[v] = 2
	}
	for _, v := range []byte("zxcvbnmZXCVBNM") {
		m[v] = 3
	}
}

func findWords(words []string) []string {
	var res []string
	for _, word := range words {
		i := 0
		l := len(word)
		for i < l && m[word[i]] == m[word[0]] {
			i++
		}
		if i == l {
			res = append(res, word)
		}
	}
	return res
}
