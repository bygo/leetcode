package main

// https://leetcode-cn.com/problems/number-of-boomerangs

func numberOfBoomerangs(points [][]int) int {
	var res int
	for _, p := range points {
		m := map[int]int{}
		for _, q := range points {
			dis := (p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1])
			m[dis]++
		}
		for _, v := range m {
			res += v * (v - 1)
		}
	}
	return res
}
