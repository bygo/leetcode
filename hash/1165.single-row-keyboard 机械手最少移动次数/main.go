package main

// https://leetcode-cn.com/problems/single-row-keyboard

func calculateTime(keyboard string, word string) int {
	var m = map[byte]int{}
	for i := range keyboard {
		m[keyboard[i]] = i
	}

	var res = m[word[0]]
	var l = len(word)
	var i = 1
	for i < l {
		res += abs(m[word[i-1]], m[word[i]])
		i++
	}

	return res
}

func abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
