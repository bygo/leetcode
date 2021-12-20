package main

// https://leetcode-cn.com/problems/4sum-ii

func fourSumCount(a, b, c, d []int) int {
	var res int
	countAB := map[int]int{}
	for i := range a {
		for j := range b {
			countAB[a[i]+b[j]]++
		}
	}
	for i := range c {
		for j := range d {
			res += countAB[-c[i]-d[j]]
		}
	}
	return res
}
