package main

import "strconv"

// https://leetcode.cn/problems/fraction-to-recurring-decimal

func fractionToDecimal(x int, y int) string {
	var res []byte
	if x*y < 0 {
		res = append(res, '-')
	}
	x = abs(x)
	y = abs(y)
	res = append(res, strconv.Itoa(x/y)...)
	x %= y
	if x == 0 {
		return string(res)
	}
	m := map[int]int{}
	for 0 < x {
		m[x] = len(res)
		x *= 10
		res = append(res, byte(x/y+'0'))
		l, ok := m[x]
		if ok {
			res = append(res, ' ')
			copy(res[l:], res[l+1:])
			res[l] = '('
			res = append(res, ')')
		}
	}
	return string(res)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
