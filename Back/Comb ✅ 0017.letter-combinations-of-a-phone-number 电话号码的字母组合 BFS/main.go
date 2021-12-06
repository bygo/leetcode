package main

// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number

func letterCombinations(digits string) []string {
	m := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}
	if len(digits) == 0 {
		return nil
	}
	queue := []string{""}
	for i := range digits {
		cnt := len(queue)
		for _, letter := range m[digits[i]] {
			for _, a := range queue[:cnt] {
				queue = append(queue, a+string(letter))
			}
		}
		queue = queue[cnt:]
	}
	return queue
}
