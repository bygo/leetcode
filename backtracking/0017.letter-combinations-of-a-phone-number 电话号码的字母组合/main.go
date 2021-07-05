package main

// Link: https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number

func letterCombinations(digits string) []string {
	var m = map[rune]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}
	if len(digits) == 0 {
		return nil
	}
	var res, combination = []string{""}, []string{}
	for _, num := range digits {
		for _, letter := range m[num] {
			for _, v := range res { //   m[num]*m[num]*""
				combination = append(combination, v+string(letter))
			}
		}
		res, combination = combination, nil
	}
	return res
}
