package main

// https://leetcode-cn.com/problems/maximum-number-of-words-you-can-type

func canBeTypedWords(text string, brokenLetters string) int {
	var m = map[byte]bool{}
	for i := range brokenLetters {
		m[brokenLetters[i]] = true
	}
	var res int
	var normal = false

	var i int
	var l = len(text)
	for i < l {
		normal = true
		for i < l && text[i] != ' ' {
			if m[text[i]] {
				normal = false
				for i < l && text[i] != ' ' {
					i++
				}
			} else {
				i++
			}
		}
		if normal {
			res++
		}
		i++
	}
	return res
}
