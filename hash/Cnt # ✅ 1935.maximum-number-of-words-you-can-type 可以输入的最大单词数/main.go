package main

// https://leetcode-cn.com/problems/maximum-number-of-words-you-can-type

func canBeTypedWords(text string, brokenLetters string) int {
	m := map[byte]bool{}
	for i := range brokenLetters {
		m[brokenLetters[i]] = true
	}
	var cur []byte
	var res int
	f := func() {
		l := len(cur)
		j := 0
		for j < l && !m[cur[j]] {
			j++
		}
		if j == l {
			res++
		}
		cur = []byte{}
	}
	for i := range text {
		ch := text[i]
		if 'a' <= ch && ch <= 'z' {
			cur = append(cur, ch)
		} else {
			f()
		}
	}
	f()
	return res
}
