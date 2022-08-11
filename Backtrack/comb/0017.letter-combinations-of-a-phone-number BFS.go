package main

// https://leetcode.cn/problems/letter-combinations-of-a-phone-number

func letterCombinations(digits string) []string {
	var numMpStrs = map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}

	if len(digits) == 0 {
		return nil
	}
	que := []string{""}
	for i := range digits {
		num := digits[i]
		queL := len(que)
		// multiplication
		for _, letter := range numMpStrs[num] {
			for _, str := range que[:queL] {
				que = append(que, str+string(letter))
			}
		}
		que = que[queL:]
	}
	return que
}

func letterCombinations_DFS(digits string) []string {
	var numMpStrs = map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}

	dL := len(digits)
	if dL == 0 {
		return nil
	}
	var str = make([]byte, dL)
	var strsComb []string
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == dL {
			strsComb = append(strsComb, string(str))
			return
		}
		num := digits[idx]
		for j := range numMpStrs[num] {
			ch := numMpStrs[num][j]
			str[idx] = ch
			dfs(idx + 1)
		}
	}

	dfs(0)
	return strsComb
}
