package main

// https://leetcode.cn/problems/letter-combinations-of-a-phone-number

func letterCombinations(digits string) []string {
	numMpStr := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}
	dL := len(digits)
	if dL == 0 {
		return nil
	}
	var buf = make([]byte, dL)
	var strs []string
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == dL {
			strs = append(strs, string(buf))
			return
		}
		num := digits[idx]
		str := numMpStr[num]
		for i := range str {
			buf[idx] = str[i]
			dfs(idx + 1)
		}
	}

	dfs(0)
	return strs
}
