package main

// https://leetcode.cn/problems/letter-combinations-of-a-phone-number

// BFS
func letterCombinations(digits string) []string {
	var digitMpStr = map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}

	dL := len(digits)
	if dL == 0 {
		return nil
	}

	strs := []string{""}
	for i := range digits {
		digit := digits[i]
		sL := len(strs)
		for _, letter := range digitMpStr[digit] { // TODO: multiplication
			for _, str := range strs[:sL] {
				strs = append(strs, str+string(letter))
			}
		}
		strs = strs[sL:]
	}
	return strs
}

// DFS
func letterCombinations(digits string) []string {
	digitMpStr := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}
	dL := len(digits)
	if dL == 0 {
		return nil
	}
	var buf = make([]byte, dL) // TODO
	var strs []string
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == dL {
			strs = append(strs, string(buf))
			return
		}
		digit := digits[idx]
		str := digitMpStr[digit]
		for i := range str {
			buf[idx] = str[i]
			dfs(idx + 1)
		}
	}

	dfs(0)
	return strs
}
