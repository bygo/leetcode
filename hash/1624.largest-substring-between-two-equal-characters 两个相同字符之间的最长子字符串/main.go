package main

// https://leetcode-cn.com/problems/largest-substring-between-two-equal-characters

func maxLengthBetweenEqualCharacters(s string) int {
	m := map[byte]int{}
	var max = -1

	for i := range s {
		_, ok := m[s[i]]
		if !ok {
			m[s[i]] = i
		} else {
			cur := i - m[s[i]] - 1
			if max < cur {
				max = cur
			}
		}
	}
	return max
}
