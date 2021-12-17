package main

// https://leetcode-cn.com/problems/find-words-that-can-be-formed-by-characters

func countCharacters(words []string, chars string) int {
	m := [26]int{}
	for i := range chars {
		m[chars[i]-'a']++
	}

	var res int
	for i := range words {
		cur := [26]int{}
		word := words[i]
		for j := range words[i] {
			ch := word[j] - 'a'
			cur[ch]++
		}
		k := 0
		for k < 26 {
			if m[k] < cur[k] {
				break
			}
			k++
		}
		if k == 26 {
			res += len(words[i])
		}
	}
	return res
}
