package main

// https://leetcode-cn.com/problems/valid-parenthesis-string

func checkValidString(s string) bool {
	var left, universal []int
	for i := range s {
		ch := s[i]
		if ch == '(' {
			left = append(left, i)
		} else if ch == '*' {
			universal = append(universal, i)
		} else if 0 < len(left) {
			left = left[:len(left)-1]
		} else if 0 < len(universal) {
			universal = universal[:len(universal)-1]
		} else {
			return false
		}
	}

	i := len(left) - 1
	j := len(universal) - 1
	for 0 <= i && 0 <= j {
		if universal[j] < left[i] {
			return false
		}
		i--
		j--
	}
	return i == -1
}
