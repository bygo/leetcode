package main

// https://leetcode-cn.com/problems/remove-vowels-from-a-string

// ❓移除元音字符

func removeVowels(s string) string {
	l := len(s)
	var res = make([]byte, 0, l)
	for i := range s {
		// 元音跳过
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			l--
		} else {
			res = append(res, s[i])
		}
	}
	return string(res[:l])
}
