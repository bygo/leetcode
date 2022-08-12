package main

// https://leetcode.cn/problems/minimum-add-to-make-parentheses-valid

func minAddToMakeValid(s string) int {
	var res int
	var left int
	for _, v := range s {
		if v == '(' {
			left++
		} else if 0 < left {
			left--
		} else {
			res++
		}
	}
	return res + left
}
