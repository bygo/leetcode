package main

// https://leetcode-cn.com/problems/find-words-that-can-be-formed-by-characters

func countCharacters(words []string, chars string) int {
	m := [26]int{}
	for i := range chars {
		m[chars[i]-'a']++
	}

	var res int
	for i := range words {
		tmp := [26]int{}
		for j := range words[i] {
			tmp[words[i][j]-'a']++
		}
		k := 0
		for ; k < 26; k++ {
			if m[k] < tmp[k] {
				break
			}
		}
		if k == 26 {
			res += len(words[i])
		}
	}
	return res
}
