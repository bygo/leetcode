package main

// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number

func letterCombinations(digits string) []string {
	m := map[rune]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}
	if len(digits) == 0 {
		return nil
	}
	res := []string{""}
	var combination []string
	for _, v := range digits {
		for _, letter := range m[v] {
			for _, a := range res {
				combination = append(combination, a+string(letter))
			}
		}
		res = combination
		combination = nil
	}
	return res
}

var m = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combinations []string


func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	back(digits, 0, "")
	return combinations
}

func back(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		digit := string(digits[index])
		word := m[digit]
		cnt := len(word)
		for i := 0; i < cnt; i++ {
			back(digits, index+1, combination+string(word[i]))
		}
	}
}
