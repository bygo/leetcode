package main

//Title: Remove Vowels from a String
// https://leetcode-cn.com/problems/remove-vowels-from-a-string

func removeVowels(s string) string {
	var res []rune
	for _, v := range s {
		if v != 'a' && v != 'e' && v != 'i' && v != 'o' && v != 'u' {
			res = append(res, v)
		}
	}
	return string(res)
}
