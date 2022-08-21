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
		for _, letter := range digitMpStr[digit] { // multiplication
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
	// equal length
	var buf = make([]byte, dL)
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

// 010110
// 101010
// 110100
// 111000

// 0100110
// 1001010
// 1010100
// 1101000
// 1110000

// 1000010110101
// 1000101011010
// 1001010101100
// 1010101010100
// 1101010101000
// 1110101010000
// 1111010100000
// 1111101000000
// 1111110000000
