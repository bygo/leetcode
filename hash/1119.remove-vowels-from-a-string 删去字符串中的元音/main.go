package main

// https://leetcode-cn.com/problems/remove-vowels-from-a-string

// :29
func removeVowels(s string) string {
	var res []byte
	for i := range s {
		if s[i] != 'a' && s[i] != 'e' && s[i] != 'i' && s[i] != 'o' && s[i] != 'u' {
			res = append(res, s[i])
		}
	}
	return string(res)
}
