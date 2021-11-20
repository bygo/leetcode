package main

// https://leetcode-cn.com/problems/second-largest-digit-in-a-string

func secondHighest(s string) int {
	var first, second = -1, -1
	for i := range s {
		if '0' <= s[i] && s[i] <= '9' {
			num := int(s[i] - '0')
			if first < num {
				second = first
				first = num
			} else if num < first && second < num {
				second = num
			}
		}
	}
	return second
}
