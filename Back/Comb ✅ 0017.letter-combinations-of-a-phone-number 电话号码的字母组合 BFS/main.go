package main

// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number

func letterCombinations(digits string) []string {
	numMpStr := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}
	if len(digits) == 0 {
		return nil
	}
	que := []string{""}
	for i := range digits {
		num := digits[i]
		queL := len(que)
		for _, letter := range numMpStr[num] {
			for _, str := range que[:queL] {
				que = append(que, str+string(letter))
			}
		}
		que = que[queL:]
	}
	return que
}
