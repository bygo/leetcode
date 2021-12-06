package main

// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number

func letterCombinations(digits string) []string {
	m := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}
	l := len(digits)
	if l == 0 {
		return []string{}
	}
	var cur = make([]byte, l)
	var res []string
	var dfs func(i int)
	dfs = func(i int) {
		if i == l {
			res = append(res, string(cur))
			return
		}
		for j := range m[digits[i]] {
			cur[i] = m[digits[i]][j]
			dfs(i + 1)
		}
	}
	dfs(0)
	return res
}
