package main

// https://leetcode-cn.com/problems/buddy-strings

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	if s == goal {
		m := [26]int{}
		for i := range s {
			m[s[i]-'a']++
		}
		for i := range m {
			if 1 < m[i] {
				return true
			}
		}
	}

	var left, right = -1, -1
	for i := range s {
		if s[i] != goal[i] {
			if left == -1 {
				left = i
			} else if right == -1 {
				right = i
			} else {
				return false
			}
		}
	}
	return right != -1 && s[left] == goal[right] && s[right] == goal[left]
}
