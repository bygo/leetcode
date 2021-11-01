package main

// https://leetcode-cn.com/problems/largest-substring-between-two-equal-characters

func maxLengthBetweenEqualCharacters(s string) int {
	m := map[byte]int{}
	var max = -1

	for r := range s {
		l, ok := m[s[r]]
		if !ok {
			m[s[r]] = r + 1
		} else {
			cur := r - l
			if max < cur {
				max = cur
			}
		}
	}
	return max
}
