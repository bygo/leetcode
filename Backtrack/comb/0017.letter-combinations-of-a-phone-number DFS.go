package main

// https://leetcode.cn/problems/letter-combinations-of-a-phone-number

func letterCombinations(digits string) []string {
	numMpStr := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}
	digitsL := len(digits)
	if digitsL == 0 {
		return nil
	}
	var str = make([]byte, digitsL)
	var strsComb []string
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == digitsL {
			strsComb = append(strsComb, string(str))
			return
		}
		num := digits[idx]
		for j := range numMpStr[num] {
			ch := numMpStr[num][j]
			str[idx] = ch
			dfs(idx + 1)
		}
	}

	dfs(0)
	return strsComb
}
