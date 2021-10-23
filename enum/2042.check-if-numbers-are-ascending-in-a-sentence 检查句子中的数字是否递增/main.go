package main

// https://leetcode-cn.com/problems/check-if-numbers-are-ascending-in-a-sentence


func areNumbersAscending(s string) bool {
	var sum = -1
	var pre = -1
	var cur int
	var num int
	for i := range s {
		num = int(s[i] - '0')
		if 0 <= num && num <= 9 {
			if sum == -1 {
				sum = num
			} else {
				sum = sum*10 + num
			}
		} else if 0 < sum {
			cur = sum
			sum = -1
			if cur <= pre {
				return false
			}
			pre = cur
		}
	}

	if -1 < sum {
		cur = sum
		if cur <= pre {
			return false
		}
		pre = cur
	}
	return true
}
