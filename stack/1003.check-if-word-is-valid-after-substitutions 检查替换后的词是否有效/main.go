package main

//Link: https://leetcode-cn.com/problems/check-if-word-is-valid-after-substitutions

func isValid(s string) bool {
	var stack []byte
	for i := range s {
		l := len(stack)
		if s[i] == 'c' && 1 < l && stack[l-2] == 'a' && stack[l-1] == 'b' {
			stack = stack[:l-2]
		} else {
			stack = append(stack, s[i])
		}
	}
	return 0 == len(stack)
}
