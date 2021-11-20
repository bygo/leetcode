package main

// https://leetcode-cn.com/problems/single-row-keyboard

func calculateTime(keyboard string, word string) int {
	var m = [26]int{}
	for i := range keyboard {
		m[keyboard[i]-'a'] = i
	}

	var res = m[word[0]-'a']
	var l = len(word)
	var i = 1
	for i < l {
		res += abs(m[word[i-1]-'a'], m[word[i]-'a'])
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
