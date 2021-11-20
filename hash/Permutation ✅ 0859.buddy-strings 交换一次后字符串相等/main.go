package main

// https://leetcode-cn.com/problems/buddy-strings

func buddyStrings(s string, goal string) bool {
	l1 := len(s)
	l2 := len(goal)
	if l1 != l2 {
		return false
	}

	if s == goal {
		m := map[byte]int{}
		for i := 0; i < l1; i++ {
			if m[s[i]] == 1 {
				return true
			}
			m[s[i]]++
		}
		return false
	}

	var first, second = -1, -1
	var cnt = [26]int{}
	for i := 0; i < l1; i++ {
		cnt[s[i]-'a']++
		if s[i] != goal[i] {
			if first == -1 {
				first = i
			} else if second == -1 {
				second = i
			} else {
				return false
			}
		}
	}

	return second != -1 && s[first] == goal[second] && s[second] == goal[first]
}
