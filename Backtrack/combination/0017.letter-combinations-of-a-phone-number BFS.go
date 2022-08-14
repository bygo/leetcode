package main

// https://leetcode.cn/problems/letter-combinations-of-a-phone-number

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
