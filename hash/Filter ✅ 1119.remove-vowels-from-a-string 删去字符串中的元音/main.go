package main

// https://leetcode-cn.com/problems/remove-vowels-from-a-string

func removeVowels(s string) string {
	l := len(s)
	var res = make([]byte, 0, l)
	for i := range s {
		if s[i] != 'a' && s[i] != 'e' && s[i] != 'i' && s[i] != 'o' && s[i] != 'u' {
			res = append(res, s[i])
		} else {
			l--
		}
	}
	return string(res[:l])
}
