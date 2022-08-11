package main

// https://leetcode.cn/problems/letter-combinations-of-a-phone-number

func letterCombinations(digits string) []string {
	var numMpStr = map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}

	if len(digits) == 0 {
		return nil
	}
	strs := []string{""}
	for i := range digits {
		num := digits[i]
		sL := len(strs)
		// multiplication
		for _, letter := range numMpStr[num] {
			for _, str := range strs[:sL] {
				strs = append(strs, str+string(letter))
			}
		}
		strs = strs[sL:]
	}
	return strs
}
