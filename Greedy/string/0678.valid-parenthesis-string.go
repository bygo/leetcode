package string

// https://leetcode.cn/problems/valid-parenthesis-string

func checkValidString(s string) bool {
	min, max := 0, 0
	for _, ch := range s {
		if ch == '(' {
			min++
			max++
		} else if ch == ')' {
			if 0 < min {
				min--
			}
			max--
			if max < 0 {
				return false
			}
			// TODO 贪心
		} else if ch == '*' {
			if 0 < min {
				min--
			}
			max++
		}
	}
	return min == 0
}
