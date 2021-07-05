package main

// Link: https://leetcode-cn.com/problems/keyboard-row

func findWords(words []string) []string {
	var res []string
	var m = map[byte]int{}
	for _, v := range []byte("qwertyuiopQWERTYUIOP") {
		m[v] = 1
	}
	for _, v := range []byte("asdfghjklASDFGHJKL") {
		m[v] = 2
	}
	for _, v := range []byte("zxcvbnmZXCVBNM") {
		m[v] = 3
	}

	for _, word := range words {
		b := true
		for i := range word {
			if m[word[i]] != m[word[0]] {
				b = false
				break
			}
		}
		if b {
			res = append(res, word)
		}
	}
	return res
}
