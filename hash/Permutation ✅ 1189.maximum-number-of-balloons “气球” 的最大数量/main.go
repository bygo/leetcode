package main

// https://leetcode-cn.com/problems/maximum-number-of-balloons

func maxNumberOfBalloons(text string) int {
	m := map[byte]int{}
	for i := range text {
		m[text[i]]++
	}
	c := map[byte]int{
		'a': 1,
		'b': 1,
		'n': 1,
		'o': 2,
		'l': 2,
	}

	var res = 1<<63 - 1
	for b := range c {
		cur := m[b] / c[b]
		if cur < res {
			res = cur
		}
	}
	return res
}
