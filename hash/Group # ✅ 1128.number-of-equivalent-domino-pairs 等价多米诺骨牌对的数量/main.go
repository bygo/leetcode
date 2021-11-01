package main

// https://leetcode-cn.com/problems/number-of-equivalent-domino-pairs

func numEquivDominoPairs(d [][]int) int {
	m := map[int]int{}
	var res int
	for i := range d {
		if d[i][0] < d[i][1] {
			d[i][0], d[i][1] = d[i][1], d[i][0]
		}
		v := d[i][0]*10 + d[i][1]
		res += m[v]
		m[v]++
	}
	return res
}
